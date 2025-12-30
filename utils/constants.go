package utils

const specialityDbName = "speciality_db"

var googleSheetsConfig = googleSheets{
	Levels:             []string{"Рівень вищої освіти", "levels"},
	Departments:        []string{"Кафедри", "departments"},
	Specialties:        []string{"Спеціальності", "specialties"},
	Programs:           []string{"Освітні програми", "programs"},
	DepartmentPrograms: []string{"Освітні програми кафедр", "department_programs"},
}
