package static

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "favorites.json",
		FileModTime: time.Unix(1605653988, 0),

		Content: string("{\n  \"sol\": [\n    {\n      \"label\": \"SOL - Stream Serum Instruction History\",\n      \"query\": \"subscription($account: String!){\\n  serumInstructionHistory(account: $account){\\n    instruction {\\n      __typename\\n      ...on  SerumNewOrder {\\n        side\\n        limitPrice\\n        maxQuantity\\n        orderType\\n        clientID\\n        accounts {\\n          market {\\n            ...AccountFragment\\n          }\\n          openOrders {\\n            ...AccountFragment\\n          }\\n          requestQueue {\\n            ...AccountFragment\\n          }\\n          payer {\\n            ...AccountFragment\\n          }\\n          owner {\\n            ...AccountFragment\\n          }\\n          coinVault {\\n            ...AccountFragment\\n          }\\n          pcVault {\\n            ...AccountFragment\\n          }\\n          splTokenProgram {\\n            ...AccountFragment\\n          }\\n          rent {\\n            ...AccountFragment\\n          }\\n          srmDiscount {\\n            ...AccountFragment\\n          }           \\n        }\\n      }\\n      ...on SerumMatchOrder{\\n        limit\\n        accounts {\\n          market {\\n            ...AccountFragment\\n          }\\n          requestQueue {\\n            ...AccountFragment\\n          }\\n          eventQueue {\\n            ...AccountFragment\\n          }\\n          bids {\\n            ...AccountFragment\\n          }\\n          asks {\\n            ...AccountFragment\\n          }\\n          coinFeeReceivable {\\n            ...AccountFragment\\n          }\\n          pcFeeReceivable {\\n            ...AccountFragment\\n          }\\n        }\\n      }\\n      ...on SerumCancelOrder{\\n        side\\n        orderId\\n        openOrders\\n        openOrderSlot\\n        accounts{\\n          market {\\n            ...AccountFragment\\n          }\\n          requestQueue {\\n            ...AccountFragment\\n          }\\n          owner {\\n            ...AccountFragment\\n          }\\n        }\\n      }\\n      ...on SerumSettleFunds {\\n        __typename\\n        accounts{\\n          market {\\n            ...AccountFragment\\n          }\\n          openOrders {\\n            ...AccountFragment\\n          }\\n          owner {\\n            ...AccountFragment\\n          }\\n          coinVault {\\n            ...AccountFragment\\n          }\\n          pcVault {\\n            ...AccountFragment\\n          }\\n          pcWallet {\\n            ...AccountFragment\\n          }\\n          signer {\\n            ...AccountFragment\\n          }\\n          splTokenProgram {\\n            ...AccountFragment\\n          }                    \\n        }\\n      }\\n      ...on SerumCancelOrderByClientId{\\n        clientID\\n        accounts{\\n          market {\\n            ...AccountFragment\\n          }\\n          openOrders {\\n            ...AccountFragment\\n          }\\n          requestQueue {\\n            ...AccountFragment\\n          }\\n          owner {\\n            ...AccountFragment\\n          }\\n        }\\n      }\\n    }\\n  }\\n}\\nfragment AccountFragment on Account {\\n  publicKey\\n  isSigner\\n  isWritable\\n}\",\n      \"variables\": {\n        \"mainnet\": \"{\\n  \\\"account\\\": \\\"YOUR-ACCOUNT-HERE\\\"\\n}\",\n        \"local\": \"{\\n  \\\"account\\\": \\\"YOUR-ACCOUNT-HERE\\\"\\n}\"\n      }\n    }\n  ]\n}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "graphiql.html",
		FileModTime: time.Unix(1605653988, 0),

		Content: string("<!DOCTYPE html>\n<!--\n Copyright 2019 dfuse Platform Inc.\n\n Licensed under the Apache License, Version 2.0 (the \"License\");\n you may not use this file except in compliance with the License.\n You may obtain a copy of the License at\n\n      http://www.apache.org/licenses/LICENSE-2.0\n\n Unless required by applicable law or agreed to in writing, software\n distributed under the License is distributed on an \"AS IS\" BASIS,\n WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n See the License for the specific language governing permissions and\n limitations under the License.\n-->\n\n<html>\n<head>\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/graphiql@0.16.0/graphiql.css\"/>\n    <link rel=\"stylesheet\" href=\"graphiql_dfuse_override.css\"/>\n    <link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Lato&display=swap\">\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js\"></script>\n    <script src=\"//unpkg.com/graphiql@0.16.0/graphiql.js\"></script>\n    <script src=\"//unpkg.com/subscriptions-transport-ws@0.8.3/browser/client.js\"></script>\n    <script src=\"//unpkg.com/graphiql-subscriptions-fetcher@0.0.2/browser/client.js\"></script>\n    <script src=\"helper.js\"></script>\n    <script src=\"https://unpkg.com/@dfuse/client@0.3.9/dist/dfuse-client.umd.js\"></script>\n</head>\n<body style=\"width: 100%; height: 100%; margin: 0; overflow: hidden;\">\n\n<div id=\"graphiql\" style=\"height: 100vh;\">Loading...</div>\n\n<script>\n    window.DfuseClientConfig = --== json . ==--;\n\n</script>\n\n<script>\n    const url = new URL(window.location.href)\n    const urlPathSegments = url.toString().split(\"/\");\n    const queryParams = url.searchParams\n\n    const server = urlPathSegments[2];\n    const proto = urlPathSegments[0];\n    const alphaSchema = isAlphaSchemaQueryParamFound(queryParams)\n\n    async function initialize() {\n        try {\n            await reconfigureGraphiQLStorage(\"sol\", \"mainnet\", alphaSchema);\n        }catch (error) {\n            console.log(\"Unable to correctly reconfigure graphiql, continuing anyway\", error)\n        }\n        loadGraphiql()\n    }\n\n    function loadGraphiql() {\n                const subscriptionsClient = new window.SubscriptionsTransportWs.SubscriptionClient((proto == 'https:' ? 'wss://' : 'ws://') + server + '/graphql',\n                    {\n                        reconnect: true,\n                        connectionCallback: (error) => {\n                            if (error != null) {\n                                alert(error.message)\n                            }\n                        }\n                    });\n\n                let activeQuery = fetchQueryProp(queryParams)\n                let activeVariables = fetchVariablesProp(queryParams)\n\n                const graphqlFetcher = graphQLFetcherFactory(queryParams)\n                const subscriptionsFetcher = window.GraphiQLSubscriptionsFetcher.graphQLFetcher(subscriptionsClient, graphqlFetcher);\n\n                ReactDOM.render(\n                    React.createElement(GraphiQL, {\n                        fetcher: subscriptionsFetcher,\n                        query: activeQuery,\n                        variables: activeVariables,\n                        onEditQuery: function(query) {\n                            activeQuery = query || undefined\n                            pushState(url, activeQuery, activeVariables)\n                        },\n                        onEditVariables: function(variables) {\n                            activeVariables = variables || undefined\n                            pushState(url, activeQuery, activeVariables)\n                        },\n                    }),\n                    document.getElementById(\"graphiql\")\n                );\n    }\n    initialize();\n\n</script>\n</body>\n</html>\n\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "graphiql_dfuse_override.css",
		FileModTime: time.Unix(1605116754, 0),

		Content: string("/**\n * Copyright 2019 dfuse Platform Inc.\n *\n * Licensed under the Apache License, Version 2.0 (the \"License\");\n * you may not use this file except in compliance with the License.\n * You may obtain a copy of the License at\n *\n *      http://www.apache.org/licenses/LICENSE-2.0\n *\n * Unless required by applicable law or agreed to in writing, software\n * distributed under the License is distributed on an \"AS IS\" BASIS,\n * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n * See the License for the specific language governing permissions and\n * limitations under the License.\n */\n\n.graphiql-container .topBar {\n    background-image: url('https://www.dfuse.io/hubfs/dfuse-graphiQL-logo-white-03-01.svg');\n    background-size: auto 31px;\n    background-repeat: no-repeat;\n    background-position: 10px 12px;\n    background-position-x: right;\n    height:44px;\n    border-bottom: 1px solid #c9cacf;\n}\n\n.graphiql-container .docExplorerShow, .graphiql-container .historyShow {\n    border-bottom: 1px solid #c9cacf!important;\n}\n\n.CodeMirror-linenumber {\n    color:#9a9ba3;\n}\n\n.graphiql-container .execute-options > li.selected, .graphiql-container .toolbar-menu-items > li.hover, .graphiql-container .toolbar-menu-items > li:active, .graphiql-container .toolbar-menu-items > li:hover, .graphiql-container .toolbar-select-options > li.hover, .graphiql-container .toolbar-select-options > li:active, .graphiql-container .toolbar-select-options > li:hover, .graphiql-container .history-contents > p:hover, .graphiql-container .history-contents > p:active {\n    background: #ff4660;\n    color: #fff;\n}\n\n.graphiql-container .topBarWrap {\n    background-image: linear-gradient(to right, rgb(255, 70, 96) 8%, rgb(65, 17, 160) 93%);\n    background-repeat: no-repeat;\n    background-position: 0px 0px;\n    background-size: cover;\n}\n\n.graphiql-container .topBar .title {\n    display:none;\n}\n\n.graphiql-container .execute-button {\n    background: rgba(255,255,255,.3);\n    border-radius: 17px;\n    border: 0px solid rgba(0, 0, 0, 0.25);\n    -webkit-box-shadow: none !important;\n    box-shadow: none !important;\n    cursor: pointer;\n    fill: #fff;\n    height: 34px;\n    margin: 0;\n    padding: 0;\n    width: 34px;\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n}\n\n.graphiql-container .execute-button svg {\n    position:relative;\n    top:1px;\n    left:1px;\n}\n\n.graphiql-container .execute-button-wrap {\n    margin: 0px 14px 0 7px;\n\n}\n\n\n.graphiql-container .toolbar-button {\n    background: rgba(255,255,255,.3);\n    -webkit-box-shadow: none !important;\n    box-shadow: none !important;\n    border-radius: 5px;\n    color: #fff;\n    cursor: pointer;\n    display: inline-block;\n    margin: 0 5px;\n    padding: 8px 20px 8px;\n    text-decoration: none;\n    text-overflow: ellipsis;\n    white-space: nowrap;\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n}\n\n.docExplorerShow,\n.history-title,\n.doc-explorer-title {\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n    letter-spacing: 1px;\n}\n\n.graphiql-container .toolbar-button:hover,\n.graphiql-container .execute-button:hover {\n    background: rgba(255,255,255,.4);\n}\n\n.graphiql-container .docExplorerShow, .graphiql-container .historyShow {\n    background: #fff;\n\n    border-bottom: 1px solid #d0d0d0;\n    border-right: none;\n    border-top: none;\n    color: #4111A0;\n    cursor: pointer;\n    font-size: 14px;\n    margin: 0;\n    outline: 0;\n    padding: 2px 20px 0 18px;\n}\n\n.graphiql-container .docExplorerShow:before {\n    border-left: 2px solid #4111A0;\n    border-top: 2px solid #4111A0;\n}\n\n.graphiql-container .doc-explorer-title-bar, .graphiql-container .history-title-bar {\n    height:44px;\n    place-items: center;\n}\n\n.graphiql-container .doc-explorer-contents, .graphiql-container .history-contents {\n    top: 57px;\n\n}\n\n.graphiql-container, .graphiql-container button, .graphiql-container input {\n    color: #22244b;\n}\n\n\n.graphiql-container .search-box {\n    border-bottom: none;\n    display: block;\n    font-size: 14px;\n    margin: 10px 0px 32px 0px;\n    position: relative;\n}\n\n.graphiql-container .search-box:before {\n    top: 4px;\n    left: 13px;\n}\n.graphiql-container .search-box > input {\n    -webkit-appearance: none;\n    -moz-appearance: none;\n    box-shadow: none !important;\n    outline: none;\n    background: #f8f8fa;\n    border: 1px solid #e1e1e4 !important;\n    padding: 13px 20px 13px 40px;\n    box-sizing: border-box;\n    margin: 0px 0px;\n    width: 100%;\n    outline: none;\n    color: #586090 !important;\n    border-radius: 8px;\n}\n\n.CodeMirror-scroll {\n    background: #f8f8fa;\n}\n\n.graphiql-container .variable-editor-title {\n    background: #e7e7ec;\n    border-top:  1px solid #dbdbe8;\n    border-bottom:  1px solid #dbdbe8;\n}\n\n\n.graphiql-container .result-window .CodeMirror-gutters, .CodeMirror-gutters {\n    background: #f0f0f5;\n    border-color: #dbdbe8;\n}\n\n.cm-keyword {\n    color: #030b8c;\n}\n\n.cm-punctuation {\n    color: #a2a2a2;\n}\n\n.cm-variable {\n    color: #4bb310;\n}\n\n.cm-atom {\n    color: #b9f;\n}\n\n.cm-property {\n    color: #4f6bc1;\n}\n\n.cm-number {\n    color: #5bcfff;\n}\n\n.cm-attribute {\n    color: #8B2BB9;\n}\n\n.cm-def {\n    color: #e60a0a;\n}\n\n.cm-string {\n    color: #E48F32;\n}\n\n.graphiql-container .type-name {\n    color: #0db4de;\n}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "helper.js",
		FileModTime: time.Unix(1605116754, 0),

		Content: string("/**\n * Copyright 2019 dfuse Platform Inc.\n *\n * Licensed under the Apache License, Version 2.0 (the \"License\");\n * you may not use this file except in compliance with the License.\n * You may obtain a copy of the License at\n *\n *      http://www.apache.org/licenses/LICENSE-2.0\n *\n * Unless required by applicable law or agreed to in writing, software\n * distributed under the License is distributed on an \"AS IS\" BASIS,\n * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n * See the License for the specific language governing permissions and\n * limitations under the License.\n */\n\nfunction refreshToken(server, token, apiKey, onCompletion, onError) {\n    if(token === \"\") {\n        return getToken(apiKey, onCompletion, onError)\n    }\n\n    const jwt = parseJwt(token);\n    const exp = jwt[\"exp\"];\n    const now = Date.now() / 1000;\n\n    console.log(\"exp  : \" + exp);\n    console.log(\"now  : \" + now);\n    const remainingTime = exp - now;\n\n    console.log(\"Time remaining in second: \" + remainingTime);\n    if (remainingTime < 60 * 60) {\n        return getToken(apiKey, onCompletion, onError)\n    }\n\n    onCompletion(token);\n}\n\nfunction getToken(apiKey, onCompletion, onError) {\n    const url = \"https://auth.dfuse.io/v1/auth/issue\";\n    const r = new XMLHttpRequest();\n    r.open(\"POST\", url, false);\n    r.setRequestHeader(\"Content-type\", \"application/json\");\n    r.onreadystatechange = function() {//Call a function when the state changes.\n        if(r.readyState === 4) {\n            if (r.status === 200) {\n                console.log(\"Got new token: \" + r.response);\n                const responseToken =  JSON.parse(r.response)\n                onCompletion(responseToken[\"token\"]);\n            } else {\n                alert(\"Error: \" + r.status + \" - \" + r.response);\n                onError(r.status, r.response);\n            }\n        }\n    };\n\n    r.send('{\"api_key\":\"' + apiKey + '\"}');\n}\n\nfunction graphQLFetcherFactory(urlQueryParams) {\n    let graphqlUrl = \"/graphql\"\n    if (isAlphaSchemaQueryParamFound(urlQueryParams)) {\n        console.log(\"Requesting alpha endpoints\")\n        graphqlUrl += \"?alpha-schema=true\"\n    }\n\n    return (graphQLParams) => {\n        return fetch(graphqlUrl, {\n            method: \"post\",\n//            headers: {\n//                Authorization: \"Bearer \" + token,\n//            },\n            body: JSON.stringify(graphQLParams),\n            credentials: \"include\",\n        }).then(function (response) {\n            return response.text();\n        }).then(function (responseBody) {\n            try {\n                return JSON.parse(responseBody);\n            } catch (error) {\n                return error.message;\n            }\n        }).catch(function (error) {\n            console.log(\"Error:\", error);\n            alert(error)\n        });\n    }\n}\n\nfunction isAlphaSchemaQueryParamFound(urlQueryParams) {\n    if (urlQueryParams.has(\"alpha-schema\")) {\n        return urlQueryParams.get(\"alpha-schema\") === \"true\"\n    }\n\n    if (urlQueryParams.has(\"alphaSchema\")) {\n        return urlQueryParams.get(\"alphaSchema\") === \"true\"\n    }\n\n    return false\n}\n\nfunction fetchQueryProp(queryParams) {\n    if (!queryParams.has(\"query\")) {\n        return undefined\n    }\n\n    try {\n        return window.atob(queryParams.get(\"query\"))\n    } catch (error) {\n        console.error(\"query params 'query' is not a valid base64 object\")\n        return undefined\n    }\n}\n\nfunction fetchVariablesProp(queryParams) {\n    if (!queryParams.has(\"variables\")) {\n        return undefined\n    }\n\n    try {\n        return window.atob(queryParams.get(\"variables\"))\n    } catch (error) {\n        console.error(\"query params 'variables' is not a valid base64 object\")\n        return undefined\n    }\n}\n\nfunction pushState(url, query, variables) {\n    const queryParams = []\n    if (query !== undefined) {\n        queryParams.push(`query=${window.btoa(query)}`)\n    }\n\n    if (variables !== undefined) {\n        queryParams.push(`variables=${window.btoa(variables)}`)\n    }\n\n    if (queryParams.length <= 0) {\n        return\n    }\n\n    window.history.pushState(\"\", \"New Query\", `${url.pathname}?${queryParams.join(\"&\")}`);\n}\n\nasync function getConfig() {\n    if (window.location.hostname === \"localhost\") {\n        return await fetchConfig()\n    }\n\n    const parts = window.location.host.split(\".\");\n    return { network: parts[0], protocol: parts[1] }\n}\n\nfunction fetchFavorites() {\n    console.info(\"Fetching favorites JSON data\")\n    return fetch(\"/graphiql/favorites.json\")\n            .then((response) => response.json())\n            .then((body) => {\n                console.log(\"Got favorites JSON data\")\n                return body\n            })\n            .catch((error) => {\n                console.log(\"Fetch favorites JSON data error\", error);\n                return {}\n            })\n}\n\nfunction fetchConfig() {\n    console.info(\"Fetching config JSON data\")\n    return fetch(\"/graphiql/config.json\")\n            .then((response) => response.json())\n            .then((body) => {\n                console.log(\"Got config JSON data\")\n                return body\n            })\n            .catch((error) => {\n                console.log(\"Fetch config JSON data error\", error);\n                return {}\n            })\n}\n\nfunction getFavoriteFromStorage() {\n    const storageItem = localStorage.getItem(\"graphiql:favorites\");\n    console.log(\"Retrieved client favorites from browser storage\")\n\n    store = { favorites: [] };\n    if (storageItem !== null) {\n        store = JSON.parse(storageItem);\n    }\n\n    return store\n}\n\nfunction setFavoriteFromStorage(store) {\n    console.log(\"Saving client favorites to browser storage\")\n    localStorage.setItem(\"graphiql:favorites\", toJSON(store));\n}\n\nasync function reconfigureGraphiQLStorage(protocol, network, alphaSchema) {\n    const isFirstTime = localStorage.getItem(\"dfuse:graphiql:is_first_time\");\n    if (isFirstTime == null) {\n        // Let's open the history pane the first time the user opens this page\n        localStorage.setItem(\"graphiql:historyPaneOpen\", toJSON(true));\n    }\n\n    const serverFavorites = await setFavorites(protocol, network, alphaSchema)\n\n    localStorage.setItem(\"dfuse:graphiql:is_first_time\", toJSON(false))\n\n    if (isFirstTime == null && serverFavorites.length > 0) {\n        localStorage.setItem(\"graphiql:query\", serverFavorites[0].query);\n\n        if (serverFavorites[0].variables) {\n            localStorage.setItem(\"graphiql:variables\", serverFavorites[0].variables);\n        }\n    }\n}\n\nasync function setFavorites(protocol, network, alphaSchema) {\n    const favoritesByProtocolMap = await fetchFavorites()\n\n    console.log(`Looking for favorites for given ${protocol}/${network} values`)\n    const serverFavorites = favoritesByProtocolMap[protocol]\n    if (serverFavorites == null) {\n        console.log(\"Favorites not found for this protocol/network values.\")\n        return\n    }\n\n    const store = getFavoriteFromStorage()\n\n    // Clear all dfuse managed favorites, we will add them back\n    store[\"favorites\"] = store[\"favorites\"].filter((value) =>\n        // We keep only favorites that don't have the `procotol`\n        value.protocol == null\n    )\n\n    console.log(\"Favorites store prior update\")\n    serverFavorites.reverse().forEach((favorite) => {\n        if (!alphaSchema && favorite.alpha) {\n            return\n        }\n\n        favorite.favorite = true\n        favorite.protocol = protocol\n        if (favorite.variables && typeof favorite.variables === \"object\") {\n            const networkVariables = favorite.variables[network]\n            if (networkVariables != null) {\n                favorite.variables = networkVariables\n            }\n        }\n\n        store[\"favorites\"] = updateFavorite(store[\"favorites\"], favorite);\n    })\n    console.log(\"Favorites store after update\")\n\n    setFavoriteFromStorage(store)\n\n    // We reverse it again because the `reverse` operation is \"in-place\"\n    return serverFavorites.reverse()\n}\n\nfunction updateFavorite(favorites, fav) {\n    const index = favorites.findIndex(f => (f.label === fav.label));\n    if (index >= 0) {\n        console.log(`Updating favorite ${fav.label}`)\n        favorites[index] = fav\n    } else {\n        console.log(`Adding favorite ${fav.label}`)\n        favorites.push(fav)\n    }\n\n    return favorites\n}\n\nfunction toJSON(input) {\n    return JSON.stringify(input)\n}\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "schema.graphql",
		FileModTime: time.Unix(1605652240, 0),

		Content: string("schema {\n  query: Queries\n  subscription: Subscription\n}\n\ntype Queries {\n}\n\ntype Subscription {\n  serumInstructionHistory(account: String!): SerumInstructionResponse\n}\n\ntype SerumInstructionResponse {\n  instruction: SerumInstruction\n}\n\n\nunion SerumInstruction = SerumInitializeMarket | SerumNewOrder | SerumMatchOrder | SerumConsumeEvents | SerumCancelOrder | SerumSettleFunds | SerumCancelOrderByClientId\n\ntype SerumInitializeMarketAccounts {\n  market: String!\n  splCoinToken: String!\n  splPriceToken: String!\n  coinMint: String!\n  priceMint: String!\n}\ntype SerumInitializeMarket {\n  baseLotSize: Uint64!\n  quoteLotSize: Uint64!\n  feeRateBps: Uint64!\n  vaultSignerNonce: Uint64!\n  quoteDustThreshold: Uint64!\n\n  accounts : SerumInitializeMarketAccounts!\n}\n\ntype Account {\n  publicKey:  String!\n  isSigner:   Boolean!\n  isWritable: Boolean!\n}\ntype SerumNewOrderAccounts {\n  market: Account!\n  openOrders: Account!\n  requestQueue: Account!\n  payer: Account!\n  owner: Account!\n  coinVault: Account!\n  pcVault: Account!\n  splTokenProgram: Account!\n  rent: Account!\n  srmDiscount: Account\n}\n\ntype SerumNewOrder {\n  side:        SideType!\n  limitPrice:  Uint64!\n  maxQuantity: Uint64!\n  orderType:   OrderType!\n  clientID:    Uint64!\n\n\n  accounts: SerumNewOrderAccounts!\n}\n\ntype SerumMatchOrderAccounts {\n  market: Account!\n  requestQueue: Account!\n  eventQueue: Account!\n  bids: Account!\n  asks: Account!\n  coinFeeReceivable: Account!\n  pcFeeReceivable: Account!\n}\ntype SerumMatchOrder {\n  limit: Uint64!\n\n  accounts : SerumMatchOrderAccounts!\n}\n\ntype SerumConsumeEventsAccounts {\n  openOrders: [Account!]!\n  market: Account!\n  eventQueue: Account!\n  coinFeeReceivable: Account!\n  pcFeeReceivable: Account!\n}\ntype SerumConsumeEvents {\n  limit: Uint64!\n\n  accounts: SerumConsumeEventsAccounts!\n}\n\ntype SerumCancelOrderAccounts {\n  market: Account!\n  requestQueue: Account!\n  owner: Account!\n}\ntype SerumCancelOrder {\n  side: SideType!\n  orderId: String!\n  openOrders: String!\n  openOrderSlot: Uint64!\n  accounts: SerumCancelOrderAccounts!\n}\n\ntype SerumSettleFundsAccounts{\n  market: Account!\n  openOrders: Account!\n  owner: Account!\n  coinVault: Account!\n  pcVault: Account!\n  coinWallet: Account!\n  pcWallet: Account!\n  signer: Account!\n  splTokenProgram: Account!\n  referrerPCWallet: Account\n}\ntype SerumSettleFunds {\n  accounts : SerumSettleFundsAccounts!\n}\n\ntype SerumCancelOrderByClientIdAccounts {\n  market: Account!\n  openOrders: Account!\n  requestQueue: Account!\n  owner: Account!\n}\ntype SerumCancelOrderByClientId {\n  clientID: Uint64!\n  accounts: SerumCancelOrderByClientIdAccounts!\n}\n\nenum SideType {\n  ASK\n  BID\n  UNKNOWN\n}\n\nenum OrderType {\n  LIMIT\n  IMMEDIATE_OR_CANCEL\n  POST_ONLY\n  UNKNOWN\n}\n\nscalar Uint64\nscalar JSON\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1605653988, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "favorites.json"
			file3, // "graphiql.html"
			file4, // "graphiql_dfuse_override.css"
			file5, // "helper.js"
			file6, // "schema.graphql"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`build`, &embedded.EmbeddedBox{
		Name: `build`,
		Time: time.Unix(1605653988, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"favorites.json":              file2,
			"graphiql.html":               file3,
			"graphiql_dfuse_override.css": file4,
			"helper.js":                   file5,
			"schema.graphql":              file6,
		},
	})
}
