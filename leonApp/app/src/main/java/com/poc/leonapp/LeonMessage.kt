package com.poc.leonapp

import org.json.JSONException
import org.json.JSONObject
import java.lang.Exception

class LeonMessage {
    var msg: String
    var action: String
    constructor(obj: JSONObject) {
        try {
            msg = obj.getString("msg")
            action = obj.getString("action")
        } catch (e: JSONException) {
            msg = "Response not valid"
            action = ""
        }
    }
    constructor(msg: String, action: String) {
        this.msg = msg
        this.action = action
    }
}