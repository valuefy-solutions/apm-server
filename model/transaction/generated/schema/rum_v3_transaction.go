// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package schema

const RUMV3Schema = `{
    "$id": "docs/spec/transactions/rumv3_transaction.json",
    "type": "object",
    "description": "An event corresponding to an incoming request or similar task occurring in a monitored service",
    "allOf": [
        {
            "properties": {
                "id": {
                    "type": "string",
                    "description": "Hex encoded 64 random bits ID of the transaction.",
                    "maxLength": 1024
                },
                "trace_id": {
                    "description": "Hex encoded 128 random bits ID of the correlated trace.",
                    "type": "string",
                    "maxLength": 1024
                },
                "parent_id": {
                    "description": "Hex encoded 64 random bits ID of the parent transaction or span. Only root transactions of a trace do not have a parent_id, otherwise it needs to be set.",
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "t": {
                    "type": "string",
                    "description": "Keyword of specific relevance in the service's domain (eg: 'request', 'backgroundjob', etc)",
                    "maxLength": 1024
                },
                "n": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "description": "Generic designation of a transaction in the scope of a single service (eg: 'GET /users/:id')",
                    "maxLength": 1024
                },
                "yc": {
                    "type": "object",
                    "properties": {
                        "sd": {
                            "type": "integer",
                            "description": "Number of correlated spans that are recorded."
                        },
                        "dd": {
                            "type": [
                                "integer",
                                "null"
                            ],
                            "description": "Number of spans that have been dd by the a recording the x."
                        }
                    },
                    "required": [
                        "sd"
                    ]
                },
                "c": {
                        "$id": "doc/spec/rum_v3_context.json",
    "title": "Context",
    "description": "Any arbitrary contextual information regarding the event, captured by the agent, optionally provided by the user",
    "type": [
        "object",
        "null"
    ],
    "properties": {
        "cu": {
            "description": "An arbitrary mapping of additional metadata to store with the event.",
            "type": [
                "object",
                "null"
            ],
            "patternProperties": {
                "^[^.*\"]*$": {}
            },
            "additionalProperties": false
        },
        "r": {
            "type": [
                "object",
                "null"
            ],
            "allOf": [
                {
                    "properties": {
                        "sc": {
                            "type": [
                                "integer",
                                "null"
                            ],
                            "description": "The status code of the http request."
                        },
                        "ts": {
                            "type": [
                                "number",
                                "null"
                            ],
                            "description": "Total size of the payload."
                        },
                        "ebs": {
                            "type": [
                                "number",
                                "null"
                            ],
                            "description": "The encoded size of the payload."
                        },
                        "dbs": {
                            "type": [
                                "number",
                                "null"
                            ],
                            "description": "The decoded size of the payload."
                        },
                        "he": {
                            "type": [
                                "object",
                                "null"
                            ],
                            "patternProperties": {
                                "[.*]*$": {
                                    "type": [
                                        "string",
                                        "array",
                                        "null"
                                    ],
                                    "items": {
                                        "type": [
                                            "string"
                                        ]
                                    }
                                }
                            }
                        }
                    }
                }
            ]
        },
        "q": {
            "properties": {
                "en": {
                    "description": "The env variable is a compounded of environment information passed from the webserver.",
                    "type": [
                        "object",
                        "null"
                    ],
                    "properties": {}
                },
                "he": {
                    "description": "Should include any headers sent by the requester. Cookies will be taken by headers if supplied.",
                    "type": [
                        "object",
                        "null"
                    ],
                    "patternProperties": {
                        "[.*]*$": {
                            "type": [
                                "string",
                                "array",
                                "null"
                            ],
                            "items": {
                                "type": [
                                    "string"
                                ]
                            }
                        }
                    }
                },
                "hve": {
                    "description": "HTTP version.",
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "mt": {
                    "description": "HTTP method.",
                    "type": "string",
                    "maxLength": 1024
                }
            },
            "required": [
                "mt"
            ]
        },
        "g": {
                "$id": "doc/spec/tags.json",
    "title": "Tags",
    "type": ["object", "null"],
    "description": "A flat mapping of user-defined tags with string, boolean or number values.",
    "patternProperties": {
        "^[^.*\"]*$": {
            "type": ["string", "boolean", "number", "null"],
            "maxLength": 1024
        }
    },
    "additionalProperties": false
        },
        "u": {
                "$id": "docs/spec/rum_v3_user.json",
    "title": "User",
    "type": [
        "object",
        "null"
    ],
    "properties": {
        "id": {
            "description": "Identifier of the logged in user, e.g. the primary key of the user",
            "type": [
                "string",
                "integer",
                "null"
            ],
            "maxLength": 1024
        },
        "em": {
            "description": "Email of the logged in user",
            "type": [
                "string",
                "null"
            ],
            "maxLength": 1024
        },
        "un": {
            "description": "The username of the logged in user",
            "type": [
                "string",
                "null"
            ],
            "maxLength": 1024
        }
    }
        },
        "p": {
            "description": "",
            "type": [
                "object",
                "null"
            ],
            "properties": {
                "rf": {
                    "description": "RUM specific field that stores the URL of the page that 'linked' to the current page.",
                    "type": [
                        "string",
                        "null"
                    ]
                },
                "url": {
                    "description": "RUM specific field that stores the URL of the current page",
                    "type": [
                        "string",
                        "null"
                    ]
                }
            }
        },
        "se": {
            "description": "Service related information can be sent per event. Provided information will override the more generic information from metadata, non provided fields will be set according to the metadata information.",
                "$id": "doc/spec/rum_v3_service.json",
    "title": "Service",
    "type": [
        "object",
        "null"
    ],
    "properties": {
        "a": {
            "description": "Name and version of the Elastic APM agent",
            "type": [
                "object",
                "null"
            ],
            "properties": {
                "n": {
                    "description": "Name of the Elastic APM agent, e.g. \"Python\"",
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "ve": {
                    "description": "Version of the Elastic APM agent, e.g.\"1.0.0\"",
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                }
            }
        },
        "fw": {
            "description": "Name and version of the web framework used",
            "type": [
                "object",
                "null"
            ],
            "properties": {
                "n": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "ve": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                }
            }
        },
        "la": {
            "description": "Name and version of the programming language used",
            "type": [
                "object",
                "null"
            ],
            "properties": {
                "n": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "ve": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                }
            }
        },
        "n": {
            "description": "Immutable name of the service emitting this event",
            "type": [
                "string",
                "null"
            ],
            "pattern": "^[a-zA-Z0-9 _-]+$",
            "maxLength": 1024
        },
        "en": {
            "description": "Environment name of the service, e.g. \"production\" or \"staging\"",
            "type": [
                "string",
                "null"
            ],
            "maxLength": 1024
        },
        "ru": {
            "description": "Name and version of the language runtime running this service",
            "type": [
                "object",
                "null"
            ],
            "properties": {
                "n": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                },
                "ve": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "maxLength": 1024
                }
            }
        },
        "ve": {
            "description": "Version of the service emitting this event",
            "type": [
                "string",
                "null"
            ],
            "maxLength": 1024
        }
    }
        }
    }
                },
                "d": {
                    "type": "number",
                    "description": "How long the transaction took to complete, in ms with 3 decimal points"
                },
                "rt": {
                    "type": [
                        "string",
                        "null"
                    ],
                    "description": "The result of the transaction. For HTTP-related transactions, this should be the status code formatted like 'HTTP 2xx'.",
                    "maxLength": 1024
                },
                "k": {
                    "type": [
                        "object",
                        "null"
                    ],
                    "description": "A mark captures the timing of a significant event during the lifetime of a transaction. Marks are organized into groups and can be set by the user or the agent.",
                    "patternProperties": {
                        "^[^.*\"]*$": {
                                "$id": "docs/spec/transactions/mark.json",
    "type": ["object", "null"],
    "description": "A mark captures the timing in milliseconds of a significant event during the lifetime of a transaction. Every mark is a simple key value pair, where the value has to be a number, and can be set by the user or the agent.",
    "patternProperties": {
        "^[^.*\"]*$": {
            "type": ["number", "null"]
        }
    },
    "additionalProperties": false
                        }
                    },
                    "additionalProperties": false
                },
                "sm": {
                    "type": [
                        "boolean",
                        "null"
                    ],
                    "description": "Transactions that are 'sampled' will include all available information. Transactions that are not sampled will not have 'spans' or 'context'. Defaults to true."
                }
            },
            "required": [
                "id",
                "trace_id",
                "yc",
                "d",
                "t"
            ]
        }
    ]
}
`
