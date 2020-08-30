package com.poc.wsleon.plugin.src

import android.content.Context
import android.content.Intent
import android.content.pm.ApplicationInfo
import android.content.pm.PackageManager
import java.util.*

class OpenApp(context: Context) {
    private val mainContext: Context = context

    fun run(appName: String) {
        val packageName: String = getPackNameByAppName(appName, mainContext)
        val launchIntent: Intent? = mainContext.packageManager.getLaunchIntentForPackage(packageName)

        if (launchIntent != null) {
            mainContext.startActivity(launchIntent)
        }
    }

    private fun getPackNameByAppName(name: String, context: Context): String {
        val pm: PackageManager = context.packageManager;
        val applicationList: List<ApplicationInfo> = pm.getInstalledApplications(PackageManager.GET_META_DATA);
        var packName = ""
        val targetAppName = name.toLowerCase(Locale.ROOT)

        for ( info: ApplicationInfo in applicationList) {
            var tmpPackName: String = pm.getApplicationLabel(info) as String

            tmpPackName = tmpPackName.toLowerCase(Locale.ROOT)
            if (tmpPackName.contains(name) || targetAppName.contains(tmpPackName)){
                packName = info.packageName;
            }
        }
        return packName
    }
}