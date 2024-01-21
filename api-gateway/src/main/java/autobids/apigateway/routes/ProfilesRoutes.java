package autobids.apigateway.routes;

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
public class ProfilesRoutes {


    @Bean
    public RouteLocator profileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/user/**", r -> r
                        .path("/profiles/user/**")
                        .uri(System.getenv("PROFILES_URI")))
                .build();
    }


    @Bean
    public RouteLocator createRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/login/me", r -> r
                        .method("GET")
                        .and()
                        .path("/profiles/login/me")
                        .filters(f -> f
                                .setPath("/profiles/user")
                                .modifyRequestBody(String.class, Profile.class, MediaType.APPLICATION_JSON_VALUE,
                                        (exchange, s) -> ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .map(authentication -> new Profile(
                                                        ((OAuth2User) authentication.getPrincipal()).getAttribute("name"),
                                                        ((OAuth2User) authentication.getPrincipal()).getAttribute("email"),
                                                        "https://media.istockphoto.com/id/1327592449/vector/default-avatar-photo-placeholder-icon-grey-profile-picture-business-man.jpg?s=612x612&w=0&k=20&c=yqoos7g9jmufJhfkbQsk-mdhKEsih6Di4WZ66t_ib7I="
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
                        .path("/profiles/edit/me")
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
                        .path("/profiles/delete/me")
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


    private static class Profile {
        private String userName;
        private String email;
        private String profileImage;

        public Profile(String userName, String email, String profileImage) {
            this.userName = userName;
            this.email = email;
            this.profileImage = profileImage;
        }

        public String getUserName() {
            return userName;
        }

        public void setUserName(String userName) {
            this.userName = userName;
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public String getProfileImage() {
            return profileImage;
        }

        public void setProfileImage(String profileImage) {
            this.profileImage = profileImage;
        }
    }
}
