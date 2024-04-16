package autobids.apigateway.routes;

import autobids.apigateway.UriConstants;
import com.fasterxml.jackson.annotation.JsonGetter;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.http.MediaType;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.security.core.context.ReactiveSecurityContextHolder;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.oauth2.core.user.OAuth2User;
import org.springframework.stereotype.Component;

import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_REQUEST_URL_ATTR;
import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.addOriginalRequestUrl;


@Component
public class Routes {
    @Bean
    public RouteLocator profileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/user/**", r -> r
                        .path(UriConstants.PROFILES_USER_ROUTES)
                        .uri(System.getenv("PROFILES_URI")))
                .build();
    }

    @Bean
    public RouteLocator adminProfilesRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/admin/profiles/**", r -> r
                        .path(UriConstants.ADMIN_ROUTES_PROFILES)
                        .uri(System.getenv("PROFILES_URI")))
                .build();
    }

    @Bean
    public RouteLocator adminCarsRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/admin/cars/**", r -> r
                        .path(UriConstants.ADMIN_ROUTES_CARS)
                        .uri(System.getenv("CARS_URI")))
                .build();
    }

    @Bean
    public RouteLocator adminMotorcyclesRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/admin/motorcycles/**", r -> r
                        .path(UriConstants.ADMIN_ROUTES_MOTORCYCLES)
                        .uri(System.getenv("MOTORCYCLES_URI")))
                .build();
    }


    @Bean
    public RouteLocator createRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/login/me", r -> r
                        .method("GET")
                        .and()
                        .path(UriConstants.PROFILES_LOGIN)
                        .filters(f -> f
                                .setPath("/profiles/user")
                                .modifyRequestBody(String.class, Profile.class, MediaType.APPLICATION_JSON_VALUE,
                                        (exchange, s) -> ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> new Profile(
                                                        ((OAuth2User) authentication.getPrincipal()).getAttribute("name"),
                                                        ((OAuth2User) authentication.getPrincipal()).getAttribute("email"),
                                                        "https://pbs.twimg.com/profile_images/1151437589062266880/AuZyoH2__400x400.jpg"
                                                ))
                                )
                        )
                        .uri(System.getenv("PROFILES_URI"))
                ).build();
    }


    @Bean
    public RouteLocator editRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/edit/me", r -> r
                        .method("PUT")
                        .and()
                        .path(UriConstants.PROFILES_EDIT)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/profiles/edit/me",
                                                            "/profiles/user/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("PROFILES_URI"))
                ).build();
    }


    @Bean
    public RouteLocator deleteSecuredProfileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/delete/me", r -> r
                        .method("DELETE")
                        .and()
                        .path(UriConstants.PROFILES_DELETE)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/profiles/delete/me",
                                                            "/profiles/delete/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("PROFILES_URI"))
                ).build();
    }

    private static class Profile {
        private String userName;
        private String email;
        private String profileImage;


        public Profile(String userName, String email, String profileImage) {
            this.userName = userName;
            this.email = email;
            this.profileImage = profileImage;
        }

        @JsonGetter("user_name")
        public String getUserName() {
            return userName;
        }

        public void setUserName(String userName) {
            this.userName = userName;
        }

        @JsonGetter("email")
        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        @JsonGetter("profile_image")
        public String getProfileImage() {
            return profileImage;
        }

        public void setProfileImage(String profileImage) {
            this.profileImage = profileImage;
        }
    }


    @Bean
    public RouteLocator getCarRoutesByPage(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/search/{page}", r -> r
                        .path(UriConstants.CARS_SEARCH)
                        .uri(System.getenv("CARS_URI")))
                .build();
    }

    @Bean
    public RouteLocator getOneCarRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/details/{id}", r -> r
                        .path(UriConstants.CARS_DETAILS)
                        .uri(System.getenv("CARS_URI")))
                .build();
    }

    @Bean
    public RouteLocator postCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/add/me", r -> r
                        .method("POST")
                        .and()
                        .path(UriConstants.CARS_ADD)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/cars/add/me",
                                                            "/cars/add/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("CARS_URI"))
                ).build();
    }

    @Bean
    public RouteLocator deleteCarRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/delete/me", r -> r
                        .method("DELETE")
                        .and()
                        .path(UriConstants.CARS_DELETE)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/cars/delete/me",
                                                            "/cars/delete/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("CARS_URI"))
                ).build();
    }

    @Bean
    public RouteLocator deleteAllCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/delete/all/me", r -> r
                        .method("DELETE")
                        .and()
                        .path(UriConstants.CARS_DELETE_ALL)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/cars/delete/all/me",
                                                            "/cars/delete/all/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("CARS_URI"))
                ).build();
    }

    @Bean
    public RouteLocator putCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/edit/me", r -> r
                        .method("PUT")
                        .and()
                        .path(UriConstants.CARS_EDIT)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/cars/edit/me",
                                                            "/cars/edit/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("CARS_URI"))
                ).build();
    }

    @Bean
    public RouteLocator getCurrentUserCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/search/user/me/{page}", r -> r
                        .method("GET")
                        .and()
                        .path(UriConstants.CARS_SEARCH_ME_ROUTES)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String[] pathSegments = path.split("/");
                                                    String page = pathSegments[pathSegments.length - 1];
                                                    String newPath = path.replaceAll(
                                                            "/cars/search/user/me/[^/]+",
                                                            "/cars/search/user/" + user.getAttribute("email")) + "/" + page;
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, request.getURI());
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("CARS_URI"))
                ).build();
    }

    @Bean
    public RouteLocator getUserCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/search/user/**", r -> r
                        .path(UriConstants.CARS_SEARCH_USER)
                        .uri(System.getenv("CARS_URI")))
                .build();
    }


    @Bean
    public RouteLocator getMotorcyclesRoutesByPage(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/search/{page}", r -> r
                        .path(UriConstants.MOTORCYCLES_SEARCH)
                        .uri(System.getenv("MOTORCYCLES_URI")))
                .build();
    }

    @Bean
    public RouteLocator getOneMotorcycleRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/details/{id}", r -> r
                        .path(UriConstants.MOTORCYCLES_DETAILS)
                        .uri(System.getenv("MOTORCYCLES_URI")))
                .build();
    }

    @Bean
    public RouteLocator postMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/add/me", r -> r
                        .method("POST")
                        .and()
                        .path(UriConstants.MOTORCYCLES_ADD)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/motorcycles/add/me",
                                                            "/motorcycles/add/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("MOTORCYCLES_URI"))
                ).build();
    }

    @Bean
    public RouteLocator deleteMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/delete/me", r -> r
                        .method("DELETE")
                        .and()
                        .path(UriConstants.MOTORCYCLES_DELETE)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/motorcycles/delete/me",
                                                            "/motorcycles/delete/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("MOTORCYCLES_URI"))
                ).build();
    }

    @Bean
    public RouteLocator deleteAllMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/delete/all/me", r -> r
                        .method("DELETE")
                        .and()
                        .path(UriConstants.MOTORCYCLES_DELETE_ALL)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/motorcycles/delete/all/me",
                                                            "/motorcycles/delete/all/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("MOTORCYCLES_URI"))
                ).build();
    }

    @Bean
    public RouteLocator putMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/edit/me", r -> r
                        .method("PUT")
                        .and()
                        .path(UriConstants.MOTORCYCLES_EDIT)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String newPath = path.replaceAll(
                                                            "/motorcycles/edit/me",
                                                            "/motorcycles/edit/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("MOTORCYCLES_URI"))
                ).build();
    }

    @Bean
    public RouteLocator getCurrentUserMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/search/user/me/{page}", r -> r
                        .method("GET")
                        .and()
                        .path(UriConstants.MOTORCYCLES_SEARCH_ME_ROUTES)
                        .filters(f -> f
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    ServerHttpRequest req = exchange.getRequest();
                                                    addOriginalRequestUrl(exchange, req.getURI());
                                                    String path = req.getURI().getRawPath();
                                                    String[] pathSegments = path.split("/");
                                                    String page = pathSegments[pathSegments.length - 1];
                                                    String newPath = path.replaceAll(
                                                            "/motorcycles/search/user/me/[^/]+",
                                                            "/motorcycles/search/user/" + user.getAttribute("email")) + "/" + page;
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, request.getURI());
                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("MOTORCYCLES_URI"))
                ).build();
    }

    @Bean
    public RouteLocator getUserMotorcyclesRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/motorcycles/search/user/**", r -> r
                        .path(UriConstants.MOTORCYCLES_SEARCH_USER)
                        .uri(System.getenv("MOTORCYCLES_URI")))
                .build();
    }

}
