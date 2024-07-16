package repo

import (
	"tekdraw-bff/model"
	"tekdraw-bff/pkg"
)

func GetDiagram(id int) (*model.Diagram, error) {
	diagram := model.Diagram{}
	err := pkg.DB.Table("diagrams").Where(
		"id = ?", id,
	).First(&diagram).Error

	if err != nil {
		return nil, err
	}

	return &diagram, nil
}

func AddDiagram(userId int, dType string, code string) (int, error) {
	diagram := model.Diagram{
		UserId: userId,
		Type:   dType,
		Code:   code,
	}
	result := pkg.DB.Table("diagrams").Create(&diagram)
	if result.Error != nil {
		return 0, result.Error
	}
	return diagram.Id, nil
}

func UpdateDiagram(id int, request model.UpdateDiagramRequest) error {
	updateData := map[string]interface{}{}
	if request.UserId != nil {
		updateData["user_id"] = *request.UserId
	}
	if request.Type != nil {
		updateData["type"] = *request.Type
	}
	if request.Code != nil {
		updateData["code"] = *request.Code
	}
	return pkg.DB.Table("diagrams").Where("id = ?", id).Updates(updateData).Error
}

func DeleteDiagram(id int) error {
	return pkg.DB.Table("diagrams").Where("id = ?", id).Delete(&model.Diagram{}).Error
}
