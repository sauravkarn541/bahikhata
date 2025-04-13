package config

// Initializes the parameters for JSON Web Tokens (JWTs)

import (
	"strings"

	"github.com/sauravkarn541/bahikhata/internal/helpers"
)

// getParamsJWT - read parameters from env
/*
	getParamsJWT() reads the parameters for JWT from the environment variables and returns
	them as an instance of the middlewares.JWTParameters struct. The function first retrieves
	the values from environment variables using os.Getenv() and then converts the values to
	the appropriate type using strconv.Atoi() where needed. The function returns the parameter
	struct and any errors encountered during the retrieval or conversion process.
*/
func getParamsJWT(env *Env) (params helpers.JWTParameters, err error) {
	params.AccessKey = []byte(strings.TrimSpace(env.AccessKey))
	params.AccessKeyTTL = env.AccessKeyTTL
	params.RefreshKey = []byte(strings.TrimSpace(env.RefreshKey))
	params.RefreshKeyTTL = env.RefreshKeyTTL

	return
}

// setParamsJWT - set parameters for JWT
/*
	setParamsJWT() sets the retrieved parameters in the jwtUtils.JWTParams struct.
	This struct is used globally throughout the project.
*/
func setParamsJWT(c helpers.JWTParameters) {
	helpers.JWTParams.AccessKey = c.AccessKey
	helpers.JWTParams.AccessKeyTTL = c.AccessKeyTTL
	helpers.JWTParams.RefreshKey = c.RefreshKey
	helpers.JWTParams.RefreshKeyTTL = c.RefreshKeyTTL
}

// InitJWTParams
/*
	InitJWTParams() initializes the JWT parameters by calling getParamsJWT() to retrieve
	the parameters, and then calls setParamsJWT() to set the parameters globally in jwtUtils.JWTParams.
*/
func InitJWTParams(env *Env) {
	var JWT helpers.JWTParameters
	JWT, err := getParamsJWT(env)
	if err != nil {
		return
	}

	// set params globally
	setParamsJWT(JWT)
}
