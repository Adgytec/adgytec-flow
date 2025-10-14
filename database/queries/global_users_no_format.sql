-- name: GetUserById :one
SELECT
	sqlc.embed(users),
	jsonb_agg_strict (
		jsonb_build_object(
			'id',
			links.id,
			'platformName',
			links.platform_name,
			'profileLink',
			links.profile_link,
			'createdAt',
			links.created_at,
			'updatedAt',
			links.updated_at
		)
	) AS social_links
FROM
	global.user_details AS users
	LEFT JOIN global.user_social_links AS links ON users.id = links.user_id
WHERE
	users.id = sqlc.arg (user_id)::UUID
GROUP BY
	users.id;
