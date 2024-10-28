package utils

import "localx/internal/models"

func ApplyTourUpdatesDetails(input, existing models.Tour) models.Tour {
	if input.Title == "" {
		input.Title = existing.Title
	}
	if input.StartTime.IsZero() {
		input.StartTime = existing.StartTime
	}
	if input.EndTime.IsZero() {
		input.EndTime = existing.EndTime
	}
	if input.GroupSize == 0 {
		input.GroupSize = existing.GroupSize
	}
	if input.Languages == "" {
		input.Languages = existing.Languages
	}
	if input.Description == "" {
		input.Description = existing.Description
	}
	if input.MeetingPlace == "" {
		input.MeetingPlace = existing.MeetingPlace
	}
	if input.ArrivalPlace == "" {
		input.ArrivalPlace = existing.ArrivalPlace
	}
	if input.WhatIsIncluded == "" {
		input.WhatIsIncluded = existing.WhatIsIncluded
	}
	if input.WhatToPrepare == "" {
		input.WhatToPrepare = existing.WhatToPrepare
	}
	if input.Prohibitions == "" {
		input.Prohibitions = existing.Prohibitions
	}
	if input.Price == 0 {
		input.Price = existing.Price
	}
	if input.Images == "" {
		input.Images = existing.Images
	}
	if input.CancellationCondition == nil {
		input.CancellationCondition = existing.CancellationCondition
	}

	// Обрабатываем булевое значение
	input.FreeCancellation = existing.FreeCancellation

	return input
}
