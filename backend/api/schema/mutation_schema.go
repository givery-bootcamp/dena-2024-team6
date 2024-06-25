package schema

type MutationSchema struct {
	// TargetID は作成・更新・削除したリソースのID
	TargetID int `json:"target_id"`
	// Message はユーザに表示するメッセージ(例：通知)
	Message string `json:"message"`
}
