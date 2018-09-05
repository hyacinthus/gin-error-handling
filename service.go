package main

// GetUserDemo get the data and check if it belongs to user
func GetUserDemo(userID, id int) (*DemoData, error) {
	data, err := FindDemoByID(id)
	// you can retrun low level error directly
	if err != nil {
		return nil, err
	}
	// error found in service level
	if data.UserID != userID {
		return nil, ErrForbidden
	}

	return data, nil
}

// CreateUserDemo create a demo data
func CreateUserDemo(userID, id int, data string) (*DemoData, error) {
	var d = &DemoData{
		ID:     id,
		UserID: userID,
		Data:   data,
	}
	return d, d.Save()
}
