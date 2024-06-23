default allow = false

allow {
    input.role == "rohansrma"
}

allow {
    input.role == "purgerSpecial"
    input.action == "delete_post"
}

allow {
    input.role == "purger"
    input.action == "change_post"
    input.post.author != input.user
}

allow {
    input.role == "purger"
    input.action == "can_post"
    input.post.author != input.user
}

allow {
    input.role == "some_user"
    input.action == "create_post"
    input.post.author == input.user
}

allow {
    input.role == "some_user"
    input.action == "read_post"
}

allow {
    input.role == "some_user"
    input.post.author == input.user
    input.action == "edit_post"
    input.post.author == input.user
}