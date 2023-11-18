package autobids.apigateway.profiles.controller;

import org.springframework.security.core.Authentication;
import org.springframework.security.oauth2.client.authentication.OAuth2AuthenticationToken;
import org.springframework.security.oauth2.core.user.OAuth2User;
import org.springframework.scheduling.annotation.Async;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.CompletableFuture;

@RestController
public class TestController {

    @GetMapping("/")
    @Async
    public CompletableFuture<String> hello() {
        return CompletableFuture.completedFuture("Hello, home");
    }

    @GetMapping("/secured")
    @Async
    public CompletableFuture<String> secured(Authentication authentication) {
        return CompletableFuture.supplyAsync(() -> {
            OAuth2AuthenticationToken oauthToken = (OAuth2AuthenticationToken) authentication;
            OAuth2User user = oauthToken.getPrincipal();
            System.out.println(user);
            String email = user.getAttribute("email");
            System.out.println("User email: " + email);
            return "Hello, secured";
        });
    }
}
