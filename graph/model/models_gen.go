// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DailyQuestion struct {
	Question string `json:"question"`
}

type NewPref struct {
	AgeRange       *string `json:"ageRange"`
	Gender         *string `json:"gender"`
	Location       *string `json:"location"`
	DarkMode       *bool   `json:"darkMode"`
	FindMatchToday *bool   `json:"findMatchToday"`
}

type NewProfile struct {
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Age        *int    `json:"age"`
	PhoneNum   *string `json:"phoneNum"`
	BirthDay   *string `json:"birthDay"`
	Gender     *string `json:"gender"`
	Bio        *string `json:"bio"`
	City       *string `json:"city"`
	ProfilePic *string `json:"profilePic"`
}

type NewUser struct {
	Profile         *NewProfile `json:"profile"`
	TodaysMatch     *string     `json:"todaysMatch"`
	LongTermMatches []*string   `json:"longTermMatches"`
	PastMatches     []*string   `json:"pastMatches"`
	LostMatches     []*string   `json:"lostMatches"`
	BlockMatches    []*string   `json:"blockMatches"`
	DailyAnswer     *string     `json:"dailyAnswer"`
	Preference      *NewPref    `json:"preference"`
}

type Pref struct {
	AgeRange       *string `json:"ageRange"`
	Gender         *string `json:"gender"`
	Location       *string `json:"location"`
	DarkMode       *bool   `json:"darkMode"`
	FindMatchToday *bool   `json:"findMatchToday"`
}

type Profile struct {
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Age        *int    `json:"age"`
	PhoneNum   *string `json:"phoneNum"`
	BirthDay   *string `json:"birthDay"`
	Gender     *string `json:"gender"`
	Bio        *string `json:"bio"`
	City       *string `json:"city"`
	ProfilePic *string `json:"profilePic"`
}

type User struct {
	UserID          string    `json:"userID"`
	Profile         *Profile  `json:"profile"`
	TodaysMatch     *string   `json:"todaysMatch"`
	LongTermMatches []*string `json:"longTermMatches"`
	LostMatches     []*string `json:"lostMatches"`
	BlockMatches    []*string `json:"blockMatches"`
	DailyAnswer     *string   `json:"dailyAnswer"`
	Preference      *Pref     `json:"preference"`
}
