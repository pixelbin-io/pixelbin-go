# 2.3.0

-   Added support for generating signed Custom Domain and PixelBin CDN urls

# 2.2.0

-   Added worker and Custom Domain url support to `UrlToObj` & `ObjToUrl`

# 2.1.0

-   Added a method for obtaining the context of a file via url.

# 2.0.0

-   Fixed `tags` being stringified inadvertently. If you are experiencing validation errors around `tags` in previous versions, you should upgrade your SDKs.
-   Fixed non-string fields that were not working, such as `Overwrite` (boolean) in FileUpload.
-   method for getting org details has changed from `getAppByToken` => `getAppOrgDetails`
