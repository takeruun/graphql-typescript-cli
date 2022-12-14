// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Connection interface {
	IsConnection()
	GetPageInfo() *PageInfo
	GetEdges() []Edge
	GetTotalCount() int
}

type Edge interface {
	IsEdge()
	GetNode() Node
	GetCursor() string
}

type Node interface {
	IsNode()
	GetID() string
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type PageCondition struct {
	First  *int    `json:"first"`
	After  *string `json:"after"`
	Last   *int    `json:"last"`
	Before *string `json:"before"`
	// 現在ページNo
	PageNo int `json:"pageNo"`
	// 1ページ表示件数
	DefautLimit *int `json:"defautLimit"`
}

type PageInfo struct {
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

func (Todo) IsNode()            {}
func (this Todo) GetID() string { return this.ID }

// ページングを伴う結果返却用
type TodoConnection struct {
	// ページ情報
	PageInfo *PageInfo `json:"pageInfo"`
	// 検索結果一覧（カーソル情報を含む）
	Edges []*TodoEdge `json:"edges"`
	// 検索結果の全件数
	TotalCount int `json:"totalCount"`
}

func (TodoConnection) IsConnection()               {}
func (this TodoConnection) GetPageInfo() *PageInfo { return this.PageInfo }
func (this TodoConnection) GetEdges() []Edge {
	if this.Edges == nil {
		return nil
	}
	interfaceSlice := make([]Edge, 0, len(this.Edges))
	for _, concrete := range this.Edges {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}
func (this TodoConnection) GetTotalCount() int { return this.TotalCount }

type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

func (TodoEdge) IsEdge()                {}
func (this TodoEdge) GetNode() Node     { return *this.Node }
func (this TodoEdge) GetCursor() string { return this.Cursor }

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos []*Todo `json:"todos"`
}
