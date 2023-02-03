# 2.0.0
* Fixed `tags` being stringified inadvertently. If you are experiencing validation errors around `tags` in previous versions, you should upgrade your SDKs.
* Fixed non-string fields that were not working, such as `Overwrite` (boolean) in FileUpload.
* method for getting org details has changed from `getAppByToken` => `getAppOrgDetails`