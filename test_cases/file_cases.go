package cases

var TestCasesHappyPath = []struct {
	Description string
	FileInput   string
	FileOutput  string
}{
	{"c1 : 2 levels", "folder_c1", "folder_c1_results"},
	{"c2 : 3 levels", "folder_c2", "folder_c2_results"},
	{"c3 : 3 levels with going back", "folder_c3", "folder_c3_results"},
	{"c5 : file without packages", "folder_c5", "folder_c5_results"},
	{"c6 : files + paths + packages", "folder_c6", "folder_c6_results"},
	{"c7 : small test with packages", "folder_c7", "folder_c7_results"},
}

var TestCasesInvalid = []struct {
	Description string
	FileInput   string
	FileOutput  string
}{
	{"invalid path", "folder_c4", "folder_c4_results"},
}
