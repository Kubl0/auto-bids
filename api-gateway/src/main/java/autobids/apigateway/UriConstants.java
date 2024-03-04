package autobids.apigateway;

public final class UriConstants {
    public static final String PROFILES_USER = "/profiles/user/**";
    public static final String PROFILES_USER_ROUTES = "/profiles/user/{email}";
    public static final String PROFILES_LOGIN = "/profiles/login/me";
    public static final String PROFILES_DELETE = "/profiles/delete/me";
    public static final String PROFILES_EDIT = "/profiles/edit/me";
    public static final String CARS_ADD = "/cars/add/me";
    public static final String CARS_DELETE = "/cars/delete/me";
    public static final String CARS_DELETE_ALL = "/cars/delete/all/me";
    public static final String CARS_EDIT = "/cars/edit/me";
    public static final String CARS_SEARCH = "/cars/search/{page}";
    public static final String CARS_DETAILS = "/cars/details/{id}";
    public static final String CARS_SEARCH_USER = "/cars/search/user/{email}/{page}";
    public static final String CARS_SEARCH_ME = "/cars/search/user/me/**";
    public static final String CARS_SEARCH_ME_ROUTES = "/cars/search/user/me/{page}";
}
