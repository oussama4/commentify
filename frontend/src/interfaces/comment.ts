import type { User } from "./user";

export interface Comment {
    body: string
    createdAt: string
    author: User
}