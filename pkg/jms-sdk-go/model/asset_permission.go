package model

import (
	"time"
)

type AssetPermission struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	IsActive    bool      `json:"is_active"`
	DateExpired time.Time `json:"date_expired"`
	CreatedBy   string    `json:"created_by"`
	DateCreated time.Time `json:"date_created"`
	Comment     string    `json:"comment"`
	DateStart   time.Time `json:"date_start"`
	OrgID       string    `json:"org_id"`
	Actions     []string  `json:"actions"`
	FromTicket  int8      `json:"from_ticket"`
}

type AssetPermissionAssetRelation struct {
	ID                int    `json:"id"`
	AssetPermissionID string `json:"assetpermission_id"`
	AssetID           string `json:"asset_id"`
}

type AssetPermissionNodeRelation struct {
	ID                int    `json:"id"`
	AssetPermissionID string `json:"assetpermission_id"`
	NodeID            string `json:"node_id"`
}

type AssetPermissionSystemUserRelation struct {
	ID                int    `json:"id"`
	AssetPermissionID string `json:"assetpermission_id"`
	SystemUserID      string `json:"systemuser_id"`
}

type AssetPermissionSystemUserGroupRelation struct {
	ID                int    `json:"id"`
	AssetPermissionID string `json:"assetpermission_id"`
	UserGroupID       string `json:"usergroup_id"`
}

type AssetPermissionUserRelation struct {
	ID                int    `json:"id"`
	AssetPermissionID string `json:"assetpermission_id"`
	UserID            string `json:"user_id"`
}