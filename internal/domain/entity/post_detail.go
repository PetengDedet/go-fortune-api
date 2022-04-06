package entity

import "gopkg.in/guregu/null.v4"

type PostDetailList struct {
	Type     null.String `json:"type"`
	OrderNum int64       `json:"order_num"`
	Value    null.String `json:"value"`
	Cover    *Cover      `json:"cover,omitempty"`
}
