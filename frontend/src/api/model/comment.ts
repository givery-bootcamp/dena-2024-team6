export interface Comment {
    id: number
    post_id: number
    user_id: number
    body: string
    created_at: string
    updated_at: string
    }

export interface CommentList {
    comments: Comment[]
}

export interface CreatComment{
    post_id: number
    user_id: number
    body: string
}

export interface UpdateComment{
    id: number
    post_id: number
    user_id: number
    body: string
}