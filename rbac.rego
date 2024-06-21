default allow = false

allow {
    input.role == "admin"
}

allow {
    input.role == "moderator"
    input.action == "create_post"
}

allow {
    input.role == "moderator"
    input.action == "edit_post"
    input.post.author != input.user
}

allow {
    input.role == "general_user"
    input.action == "create_post"
    input.post.author == input.user
}

allow {
    input.role == "general_user"
    input.action == "read_post"
}

allow {
    input.role == "general_user"
    input.post.author == input.user
    input.action == "delete_post"
    input.post.author == input.user
}

allow {
    input.role == "general_user"
    input.post.author == input.user
    input.action == "edit_post"
    input.post.author == input.user
}