package autobids.apigateway.security;

import org.springframework.http.HttpStatus;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.server.WebFilterExchange;
import org.springframework.security.web.server.authentication.ServerAuthenticationFailureHandler;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.net.URI;

@Component
public class OAuth2LoginFailureHandler implements ServerAuthenticationFailureHandler {

    @Override
    public Mono<Void> onAuthenticationFailure(WebFilterExchange webFilterExchange, AuthenticationException exception) {
        String frontendUri = System.getenv("FRONTEND_URI");
        String redirectUrl = frontendUri + "/#/register";

        return Mono.fromRunnable(() -> {
            ServerWebExchange serverWebExchange = webFilterExchange.getExchange();
            serverWebExchange.getResponse().getHeaders().setLocation(URI.create(redirectUrl));
            serverWebExchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
        });
    }
}
