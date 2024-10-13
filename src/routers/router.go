package router

import (
	"Skool_Saver/src/entities"
	"log"
	"net/http"
	"strings"
)

func NewRouter() {
	var excludedRoutes = []string{}
	for _, route := range routes {
		if route.Method == "POST" {
			http.Handle("/skoolsaver"+route.Pattern, PostMiddleware(TokenMiddleware(excludedRoutes, http.HandlerFunc(route.HandlerFunc))))
		} else {
			http.Handle("/skoolsaver"+route.Pattern, GetMiddleware(TokenMiddleware(excludedRoutes, http.HandlerFunc(route.HandlerFunc))))
		}
	}
}

func PostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// POST method checking. No other method will be allowed within this function
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type, Authorization,Auth")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With,Auth")
		w.Header().Set("Access-Control-Expose-Headers", "Autho")
		w.Header().Set("Cache-Control", "no-cache,no-store")
		if req.Method != "POST" {
			log.Println(req.Method + " is called in " + req.URL.Path)
			entities.ThrowJSONResponse(entities.NotPostMethodResponse(), w)
			return
		}
		next.ServeHTTP(w, req)
	})
}

// GetMiddleware method is used to handle post method. No other method will not applied
func GetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// POST method checking. No other method will be allowed within this function
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type, Authorization,Auth")
		w.Header().Set("Access-Control-Expose-Headers", "Autho")
		w.Header().Set("Cache-Control", "no-cache,no-store")
		if req.Method != "GET" {
			log.Println(req.Method + " is called in " + req.URL.Path)
			entities.ThrowJSONResponse(entities.NotPostMethodResponse(), w)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func TokenMiddleware(excludedRoutes []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, route := range excludedRoutes {
			if strings.HasPrefix(req.URL.Path, route) {
				// Pass the request to the next handler without JWT verification
				next.ServeHTTP(w, req)
				return
			}
		}
		// 	var token = req.Header.Get("Authorization")
		// 	if len(token) == 0 {
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		return
		// 	}
		// 	claims, success, err, errtype := utility.ValidateToken(token)
		// 	if !success {
		// 		log.Print("\n\n token error:", err, errtype)
		// 		logger.Log.Print("\n\n token error:", err, errtype)
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		return

		// 	}

		// 	isAdminClient, _, _ := models.CheckClientPermissions(claims.ClientId)
		// 	IsOrgAdmin, _ := models.CheckTheUserIsAdmin(claims.UserId)

		// 	if req.Method == "POST" && !isAdminClient {
		// 		reqData := make(map[string]interface{})
		// 		err1 := json.NewDecoder(req.Body).Decode(&reqData)
		// 		if err1 != nil {
		// 			logger.Log.Println("Error while decoding request body -------->", err1)
		// 			http.Error(w, err1.Error(), http.StatusBadRequest)
		// 			return
		// 		}
		// 		reqData["clientid"] = claims.ClientId
		// 		reqData["userid"] = claims.UserId

		// 		if !IsOrgAdmin {
		// 			ValidOrgIds, _ := models.FetchUserOrgnizationIds(claims.UserId)
		// 			validOrgId := false
		// 			if mstOrgId, ok := reqData["mstorgnhirarchyid"]; ok && mstOrgId != nil {
		// 				// Attempt to convert the value to float64
		// 				if id, ok := mstOrgId.(float64); ok {
		// 					// Convert to int64 and proceed with your logic
		// 					for i := 0; i < len(ValidOrgIds); i++ {
		// 						if ValidOrgIds[i] == int64(id) {
		// 							validOrgId = true
		// 							break
		// 						}
		// 					}
		// 				}
		// 			}
		// 			logger.Log.Println("validOrg Id: ", validOrgId)

		// 			// If the organization ID is not valid, update it
		// 			if !validOrgId {
		// 				reqData["mstorgnhirarchyid"] = claims.MstorgnhirarchyId
		// 			}
		// 			logger.Log.Println("validOrg Id2: ", validOrgId)
		// 		}
		// 		// reqData["mstorgnhirarchyid"] = claims.MstorgnhirarchyId

		// 		var buf bytes.Buffer
		// 		err2 := json.NewEncoder(&buf).Encode(reqData)
		// 		if err2 != nil {
		// 			logger.Log.Println("Error while encoding request body -------->", err2)
		// 			w.WriteHeader(http.StatusBadGateway)
		// 			return
		// 		}
		// 		req.Body = io.NopCloser(&buf)
		// 		req.ContentLength = int64(buf.Len())
		// 	} else if req.Method == "GET" && !isAdminClient {
		// 		var query = req.URL.Query()
		// 		queryClientId := query.Get("clientid")
		// 		queryMstOrgnHirarchyId := query.Get("mstorgnhirarchyid")
		// 		queryUserId := query.Get("userid")
		// 		authorized := true
		// 		if queryClientId != "" {
		// 			clientId, _ := strconv.ParseInt(queryClientId, 10, 64)
		// 			if clientId != claims.ClientId {
		// 				authorized = false
		// 			}
		// 		}
		// 		if queryMstOrgnHirarchyId != "" {
		// 			mstOrgnHirarchyId, _ := strconv.ParseInt(query.Get("mstorgnhirarchyid"), 10, 64)
		// 			if !IsOrgAdmin {
		// 				ValidOrgIds, _ := models.FetchUserOrgnizationIds(claims.UserId)
		// 				validOrgId := false
		// 				for i := 0; i < len(ValidOrgIds); i++ {
		// 					if ValidOrgIds[i] == mstOrgnHirarchyId {
		// 						validOrgId = true
		// 						break
		// 					}
		// 				}
		// 				logger.Log.Println("Valid Org:", validOrgId)
		// 				if !validOrgId {
		// 					authorized = false
		// 				}
		// 			}
		// 		}
		// 		if queryUserId != "" {
		// 			userId, _ := strconv.ParseInt(query.Get("userid"), 10, 64)
		// 			if userId != claims.UserId {
		// 				authorized = false
		// 			}
		// 		}
		// 		if !authorized {
		// 			w.WriteHeader(http.StatusUnauthorized)
		// 			return
		// 		}
		// 	}

		next.ServeHTTP(w, req)
	})
}
func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.ThrowJSONResponse(entities.BlankPathCheckResponse(), w)
}
