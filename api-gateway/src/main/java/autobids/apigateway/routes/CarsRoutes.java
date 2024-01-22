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
public class CarsRoutes {
    @Bean
    public RouteLocator carsRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/page/**", r -> r
                        .path("/cars/page/**")
                        .uri(System.getenv("CARS_URI")))
                .build();
    }

    @Bean
    public RouteLocator postCarsRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/cars/add/me", r -> r
                        .method("POST")
                        .and()
                        .path("/cars/add/me")
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
                        .path("/cars/delete/me")
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
                        .path("/cars/delete/all/me")
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
                        .path("/cars/edit/me")
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
}
