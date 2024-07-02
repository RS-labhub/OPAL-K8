# Define the regions and the minimum required Karma for each permission level
region_karma_requirements = {
    "us-east-1": {"read": 10, "write": 50, "admin": 100},
    "us-west-2": {"read": 5,  "write": 30, "admin": 90},
    "eu-central-1": {"read": 20, "write": 60, "admin": 110},
    "ap-southeast-1": {"read": 15, "write": 40, "admin": 95}
}

# Mock data source for user Karma values
user_karma = {
    "user-1": 55,
    "user-2": 45,
    "user-3": 120
}

# Mock data source for user region values
user_region = {
    "user-1": "us-east-1",
    "user-2": "us-west-2",
    "user-3": "eu-central-1"
}

# Define the input structure
# input: {
#   "user": {
#     "id": "user-1"
#   },
#   "requested_permission": "write"
# }

# Calculate whether the user has sufficient Karma for the requested permission
allow {
    user_id := input.user.id
    requested_permission := input.requested_permission
    region := user_region[user_id]
    required_karma := region_karma_requirements[region][requested_permission]
    user_karma[user_id] >= required_karma
}