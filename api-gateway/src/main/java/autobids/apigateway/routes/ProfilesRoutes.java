package autobids.apigateway.routes;

import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
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
    public RouteLocator securedProfileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/me", r -> r
                        .path("/profiles/me")
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
                                                            "/profiles/me",
                                                            "/profiles/user/" + user.getAttribute("email"));
                                                    exchange.getAttributes().put(GATEWAY_REQUEST_URL_ATTR, newPath);
                                                    ServerHttpRequest request = req.mutate().path(newPath).build();

                                                    return exchange.mutate().request(request).build();
                                                })
                                                .flatMap(chain::filter)
                                )
                        )
                        .uri(System.getenv("PROFILES_URI")))
                .build();
    }

//    @Bean
//    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
//        return builder.routes()
//                .route("/profiles/me", r -> r
//                        .method("POST")
//                        .and()
//                        .path("/profiles/me")
//                        .filters(f -> f
//                                .modifyRequestBody(
//                                        String.class,
//                                        Profile.class,
//                                        MediaType.APPLICATION_JSON_VALUE,
//                                        (exchange, s) -> ReactiveSecurityContextHolder.getContext()
//                                                .map(SecurityContext::getAuthentication)
//                                                .map(authentication -> {
//                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
//                                                    return Mono.just(new Profile(user.getName(), user.getAttribute("email"), "sampleimage.pl"));
//                                                })
//                                ).uri(System.getenv("PROFILES_URI")))
//
//                        )
//                .build();
//    }
}

