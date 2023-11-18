package autobids.apigateway.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.scheduling.annotation.Async;
import org.springframework.web.bind.annotation.*;

import java.util.concurrent.CompletableFuture;

@RestController
@RequestMapping(path = "/profile")
public class ProfilesController {

    @GetMapping("/user/email={email}")
    @Async
    public CompletableFuture<ResponseEntity<String>> get_user_profile(@PathVariable String email) {
        return CompletableFuture.completedFuture(new ResponseEntity<>(String.format("""
                {
                    "pub_id": "5f4dcc3b5aa765d61d8327deb882cf99",
                    "user_name": "john_doe",
                    "email": "%s",
                    "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                }
                """, email),HttpStatus.OK));
    }

    @GetMapping("/list/page={page}")
    @Async
    public CompletableFuture<ResponseEntity<String>> get_profiles(@PathVariable Integer page) {
        return CompletableFuture.completedFuture(new ResponseEntity<>(
                """
                [
                    {
                        "pub_id": "5f4dcc3b5aa765d61d8327deb882cf99",
                        "user_name": "john_doe",
                        "email": "john.doe@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "6d0c3e7a8e4c5f2a1b3e4f5c",
                        "user_name": "alice_smith",
                        "email": "alice.smith@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "7a3b5f1c2e4d6c8a9f1b2e3d",
                        "user_name": "bob_jones",
                        "email": "bob.jones@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "2e4d6c8a9f1b2e3d7a3b5f1c",
                        "user_name": "emma_wilson",
                        "email": "emma.wilson@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "1b2e3d7a3b5f1c2e4d6c8a9f",
                        "user_name": "charlie_brown",
                        "email": "charlie.brown@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "8e4c5f2a1b3e4f5c6d0c3e7a",
                        "user_name": "olivia_miller",
                        "email": "olivia.miller@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "9f1b2e3d7a3b5f1c2e4d6c8a",
                        "user_name": "david_wright",
                        "email": "david.wright@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "3d7a3b5f1c2e4d6c8a9f1b2e",
                        "user_name": "sophia_hall",
                        "email": "sophia.hall@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "5f1c2e4d6c8a9f1b2e3d7a3b",
                        "user_name": "michael_smith",
                        "email": "michael.smith@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    },
                    {
                        "pub_id": "c8a9f1b2e3d7a3b5f1c2e4d6",
                        "user_name": "samantha_davis",
                        "email": "samantha.davis@example.com",
                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                    }
                ]
                """, HttpStatus.OK));
    }

//    @PutMapping("/user/email={email}")
//    @Async
//    public CompletableFuture<String>
}
