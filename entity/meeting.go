package entity

//会议信息数据结构
type MeetingData struct{
	Title 		 string 	`json:"title"`
	Sponsor 	 string 	`json:"sponsor"`
	Participator []string   `json:"participator"`
	Start 		 string 	`json:"start"`
	End 		 string 	`json:"end"`
}

//得到会议发起者
func (m *MeetingData) GetSponsor() string{
	return m.Sponsor
}

//得到会议参与者
func (m *MeetingData) GetParticipator() []string{
	return m.Participator
}

//增加会议参与者
func (m *MeetingData) AddParticipator(participator string) {
	m.Participator = append(m.Participator, participator)
}

//删除会议参与者
func (m *MeetingData) RemoveParticipator(participator string) {
	for i, meeting := range m.Participator {
		if meeting == participator {
			m.Participator = append(m.Participator[:i], m.Participator[i+1:]...)
		}
	}
}

//获得开始时间
func (m *MeetingData) GetStartDate() string{
	return m.Start
}

//获得结束时间
func (m *MeetingData) GetEndDate() string{
	return m.End
}

//获得会议标题
func (m *MeetingData) GetTitle() string{
	return m.Title
}

//判断用户是否参与会议
func (m *MeetingData) IsParticipator(username string) bool{
	for _, meeting := range m.Participator {
		if meeting == username {
			return true
		}
	}
	return false
}