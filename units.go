package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// GetUnitsPath : Get unit files path
func getUnitsPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Cannot find user home path: " + err.Error())
	}
	unitsPath := path.Join(home, "/.config/systemd/user/")
	return unitsPath
}

// CreateUnitsPath : Create systemd Path
func CreateUnitsPath() {
	unitsPath := getUnitsPath()
	_, err := os.Stat(unitsPath) // Check if "$HOME/.config/systemd/user" exists
	if err != nil {
		os.MkdirAll(unitsPath, os.ModePerm)
	}
}

// disableUnit : Disable systemd unit before deleting unit files
func disableUnit(unitFile string) {
	unitName := unitFile[strings.LastIndex(unitFile, "/")+1:]
	if strings.HasSuffix(unitName, "timer") {
		cmd := exec.Command("systemctl", "--user", "disable", "--now", unitName)
		err := cmd.Run()
		if err != nil {
			log.Fatal("Cannot run external command to disable systemd timer.", err)
		}
	}
}

// DeleteAllUnits : Delete all unit files
func DeleteAllUnits() {
	files, err := filepath.Glob(path.Join(getUnitsPath(), UnitPrefix+"*"))
	if err != nil {
		log.Fatal("Cannot get list of unit files.", err)
	}
	for _, f := range files {
		disableUnit(f)
		if err := os.Remove(f); err != nil {
			log.Fatal("Cannot delete systemd files.", err)
		}
	}
}

// WriteToFile : Write string to specified file
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot write to file.", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, data)
	if err != nil {
		log.Fatal("Cannot write string to file.", err)
	}
	return file.Sync()
}

// CreateUnit : Create service and timer units
func CreateUnit(index int, record []string) {
	var err error
	var notificationLength, NotificationRepetition int
	var eventName, eventType, eventDate, eventCalendar, icon, comment string
	id := index + 1
	eventName = record[0]
	eventType = record[1]
	eventDate = record[2]
	eventCalendar = record[3]
	notificationLength, err = strconv.Atoi(record[4])
	if err != nil {
		log.Fatal("Cannot convert duration to integer.", err)
	}
	NotificationRepetition, err = strconv.Atoi(record[5])
	if err != nil {
		log.Fatal("Cannot convert repetition interval to integer.", err)
	}
	icon = record[6]
	comment = record[7]

	fileName := UnitPrefix + strconv.Itoa(id) + "_" + eventName + "_" + eventType

	filePath := path.Join(getUnitsPath(), fileName)
	// Write service file
	writeServiceErr := WriteToFile(filePath+".service",
		"[Unit]\n"+
			"Description="+eventName+" "+eventType+" "+eventDate+" (calendar: "+eventCalendar+")\n"+
			"\n"+
			"[Service]\n"+
			"ExecStart=notify-send -u normal "+"-t "+strconv.Itoa(notificationLength*1000)+" -a \"System\" -i "+icon+" \""+eventName+" "+eventType+"\" \""+comment+"\"\n"+
			"\n"+
			"[Install]\n"+
			"WantedBy=default.target\n")

	if writeServiceErr != nil {
		log.Fatal("Cannot write service unit to file: "+fileName, err)
	}

	var year, month, day int
	// Calculate Year Month Day
	if strings.ToLower(eventCalendar) == "gregorian" {
		s := strings.Split(eventDate, "-")
		year, err = strconv.Atoi(s[0])
		if err != nil {
			log.Fatal("Cannot convert year to integer: " + s[0] + ". " + err.Error())
		}
		month, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatal("Cannot convert month to integer: " + s[1] + ". " + err.Error())
		}
		day, err = strconv.Atoi(s[2])
		if err != nil {
			log.Fatal("Cannot convert day to integer: " + s[2] + ". " + err.Error())
		}
	} else if strings.ToLower(eventCalendar) == "jalali" {
		s := strings.Split(eventDate, "/")
		year, err = strconv.Atoi(s[0])
		if err != nil {
			log.Fatal("Cannot convert year to integer: " + s[0] + ". " + err.Error())
		}
		month, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatal("Cannot convert month to integer: " + s[1] + ". " + err.Error())
		}
		day, err = strconv.Atoi(s[2])
		if err != nil {
			log.Fatal("Cannot convert day to integer: " + s[2] + ". " + err.Error())
		}
		year, month, day = JalaliToGregorian(year, month, day)
	} else {
		log.Fatal("Unknown calendar type: "+eventCalendar+".", err)
	}
	// strYear := strconv.Itoa(year)
	strMonth := strconv.Itoa(month)
	strDay := strconv.Itoa(day)
	// Write timer file
	writeTimerErr := WriteToFile(filePath+".timer",
		"[Unit]\n"+
			"Description="+eventName+" "+eventType+" "+eventDate+
			" (calendar: "+eventCalendar+")\n"+
			"Requires="+fileName+".service\n"+
			"\n"+
			"[Timer]\n"+
			"Unit="+fileName+".service\n"+
			"OnCalendar=*-"+strMonth+"-"+strDay+
			" 00/"+strconv.Itoa(NotificationRepetition)+":00:00\n"+
			"\n"+
			"[Install]\n"+
			"WantedBy=timers.target\n")

	if writeTimerErr != nil {
		log.Fatal("Cannot write service unit to file: "+fileName, err)
	}

	// Enable timer
	cmd := exec.Command("systemctl", "--user", "enable", "--now", fileName+".timer")
	EnableUnitErr := cmd.Run()
	if EnableUnitErr != nil {
		log.Fatal("Cannot enable systemd unit timer.", err)
	}
}
