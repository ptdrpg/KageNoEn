package repository

import "KageNoEn/model"

func (r *Repository) GetAllFriends(userId string) ([]model.FriendWithStatus, error) {
	var friends []model.FriendWithStatus
	if err := r.DB.Table("users u").
		Select("u.id, u.username, u.is_online, f.status").
		Joins("JOIN friend_lists f ON (u.id = f.sender OR u.id = f.receiver)").
		Where("(f.sender = ? OR f.receiver = ?) AND f.status = ?", userId, userId, "accepted").
		Scan(&friends).Error; err != nil {
		return []model.FriendWithStatus{}, err
	}
	return friends, nil
}

func (r *Repository) GetFriendRequest(receiverId string) ([]string, error) {
	var requests []string
	if err := r.DB.Table("users u").
		Select("u.username").
		Joins("JOIN friend_lists f ON f.receiver = u.id").
		Where("status = ?", "pending").
		Scan(&requests).Error; err != nil {
		return []string{}, err
	}
	return requests, nil
}

func (r *Repository) AddFriend(senderId string, receiverId string) error {
	if err := r.DB.Create(&model.FriendList{
		Sender:   senderId,
		Receiver: receiverId,
		Status:   "pending",
	}).Error; err != nil {
		return err
	}
	return nil
}
