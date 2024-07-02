package autobids.apigateway.security;

import org.springframework.http.HttpStatus;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.server.WebFilterExchange;
import org.springframework.security.web.server.authentication.ServerAuthenticationSuccessHandler;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.net.URI;

@Component
public class OAuth2LoginSuccessHandler implements ServerAuthenticationSuccessHandler {

    @Override
    public Mono<Void> onAuthenticationSuccess(WebFilterExchange webFilterExchange, Authentication authentication) {
        String redirectUrl = System.getenv("FRONTEND_URI" + "/#/account");
        ServerWebExchange serverWebExchange = webFilterExchange.getExchange();
        serverWebExchange.getResponse().getHeaders().setLocation(URI.create(redirectUrl));
        serverWebExchange.getResponse().setStatusCode(HttpStatus.SEE_OTHER);

        return Mono.empty();
    }
}
