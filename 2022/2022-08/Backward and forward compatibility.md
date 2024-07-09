# Backward and forward compatibility

From https://us.mitsubishielectric.com/fa/en/resources/blog/assets/iiotcompatibility#:~:text=Backward%20compatibility%20is%20a%20design,Hardware: 
> Backward compatibility is a design that is compatible with previous versions of itself. 
> Forward compatibility is a design that is compatible with future versions of itself.

## Backward compatibility

Examples:
- omit adding new required field
- the release segregation in case of adding new field in DB on two: adding new field to DB, using such field in application

## Forward compatibility

Examples:
- tolerating to the new fields in response it receives
- tolerating to the new values for fields in response it receives
- For java has sense to do not use enums for integration related code where there is a chance to receive unexpected value for enum.

#design #backward_compatibility #forward_compatibility
#draft