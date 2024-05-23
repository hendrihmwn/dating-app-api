package pg

const LIST_USER = `
	SELECT DISTINCT users.id,
			users.name,
			users.email,
			user_profiles.image,
			user_profiles.birthdate,
			user_profiles.location,
			user_profiles.gender,
			users.created_at,
			users.updated_at,
			users.deleted_at,
			user_packages.package_act as label
	FROM users
	JOIN user_profiles ON user_profiles.user_id = users.id
	LEFT JOIN user_packages ON user_packages.user_id = users.id AND user_packages.package_act = 'verified-label' AND user_packages.expired_at > current_date
	WHERE users.deleted_at IS NULL
`

const CREATE_USER = `
	INSERT INTO users (name, email, password)
	VALUES (:name, :email, :password)
	RETURNING id, name, email, password, created_at, updated_at, deleted_at;
`

const CREATE_PROFILE = `
	INSERT INTO user_profiles (user_id, image, birthdate, gender, location)
	VALUES (:user_id, :image, :birthdate, :gender, :location)
	RETURNING id, user_id, image, birthdate, gender, location, created_at, updated_at, deleted_at;
`

const GET_USER_BY_EMAIL = `
	SELECT 	users.id,
			users.name,
			users.email,
			users.password,
			users.created_at,
			users.updated_at,
			users.deleted_at
	FROM users
	WHERE users.email = $1
	AND users.deleted_at IS NULL
	LIMIT 1;
`

const COUNT_USER_PAIR_TODAY = `
	SELECT COUNT(DISTINCT (user_pairs.pair_user_id, user_pairs.status))
	FROM user_pairs
	WHERE user_pairs.user_id = $1 
	  AND user_pairs.created_at >= current_date AND user_pairs.created_at < current_date + interval '1 day'
	  AND user_pairs.deleted_at IS NULL
`

const CREATE_USER_PAIR = `
	INSERT INTO user_pairs (user_id, pair_user_id, status)
	VALUES (:user_id, :pair_user_id, :status)
	RETURNING id, user_id, pair_user_id, status, created_at, updated_at, deleted_at;
`

const GET_PACKAGE = `
	SELECT 	packages.id,
			packages.name,
			packages.act,
			packages.price,
			packages.valid_months,
			packages.created_at,
			packages.updated_at,
			packages.deleted_at
	FROM packages
	WHERE packages.id = $1
	AND packages.deleted_at IS NULL
	LIMIT 1;
`

const LIST_PACKAGE = `
	SELECT 	packages.id,
			packages.name,
			packages.act,
			packages.price,
			packages.valid_months,
			packages.created_at,
			packages.updated_at,
			packages.deleted_at
	FROM packages
	WHERE packages.deleted_at IS NULL
`

const CREATE_USER_PACKAGE = `
	INSERT INTO user_packages (user_id, package_id, package_act, expired_at)
	VALUES (:user_id, :package_id, :package_act, :expired_at)
	RETURNING id, user_id, package_id, package_act, expired_at, created_at, updated_at, deleted_at;
`

const LIST_USER_PACKAGE = `
	SELECT 	user_packages.id,
			user_packages.user_id,
			user_packages.package_id,
			user_packages.package_act,
			user_packages.expired_at,
			user_packages.created_at,
			user_packages.updated_at,
			user_packages.deleted_at
	FROM user_packages
	WHERE user_packages.deleted_at IS NULL AND user_packages.expired_at > current_date AND user_packages.user_id = $1
`

const CREATE_ORDER = `
	INSERT INTO orders (user_id, package_name, package_price, status)
	VALUES (:user_id, :package_name, :package_price, :status)
	RETURNING id, user_id, package_name, package_price, status, created_at, updated_at, deleted_at;
`

const GET_USER_PACKAGE = `
	SELECT 	user_packages.id,
			user_packages.user_id,
			user_packages.package_id,
			user_packages.package_act,
			user_packages.expired_at,
			user_packages.created_at,
			user_packages.updated_at,
			user_packages.deleted_at
	FROM user_packages
	WHERE user_packages.user_id = $1 AND user_packages.deleted_at IS NULL AND user_packages.expired_at > current_date
	LIMIT 1;
`
