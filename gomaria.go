package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//TODO: convert to test

// connecting to maria db and saving all the recommendations subjects to the
func maina() {
	connString := "root:root@tcp(maria:3306)/"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	results, err := db.Query("SELECT *  FROM Curriculum.Subjects")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var allRecords []SubjectsCsvStruct
	var oneRecord SubjectsCsvStruct
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&oneRecord.CourseName, &oneRecord.Link, &oneRecord.Sws, &oneRecord.Language.German, &oneRecord.Language.English,
			&oneRecord.Coursetype.Komplexpraktikum, &oneRecord.Coursetype.Seminar, &oneRecord.Coursetype.Vorlesung, &oneRecord.Preferences.SOFTWARE,
			&oneRecord.Preferences.AI, &oneRecord.Preferences.LOWLEVEL, &oneRecord.Preferences.SECURITY, &oneRecord.Preferences.WEB, &oneRecord.Preferences.THEORETICAL,
			&oneRecord.Semester.Sommersemester, &oneRecord.Semester.Wintersemester)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		allRecords = append(allRecords, oneRecord)
	}
}

// connecting to maria db and saving all the recommendations subjects to the
func getAllSubjectsFromMariaDb() []SubjectsCsvStruct {
	connString := "root:root@tcp(maria:3306)/"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	results, err := db.Query("SELECT *  FROM Curriculum.Subjects")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var allRecords []SubjectsCsvStruct
	var oneRecord SubjectsCsvStruct
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&oneRecord.CourseName, &oneRecord.Link, &oneRecord.Sws, &oneRecord.Language.German, &oneRecord.Language.English,
			&oneRecord.Coursetype.Komplexpraktikum, &oneRecord.Coursetype.Seminar, &oneRecord.Coursetype.Vorlesung, &oneRecord.Preferences.SOFTWARE,
			&oneRecord.Preferences.AI, &oneRecord.Preferences.LOWLEVEL, &oneRecord.Preferences.SECURITY, &oneRecord.Preferences.WEB, &oneRecord.Preferences.THEORETICAL,
			&oneRecord.Semester.Sommersemester, &oneRecord.Semester.Wintersemester)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		allRecords = append(allRecords, oneRecord)
	}
	return allRecords
}
