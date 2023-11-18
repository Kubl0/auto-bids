package autobids.apigateway.profiles.controller;

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
        return CompletableFuture.completedFuture(
                new ResponseEntity<>(String.format(
                        """
                                {
                                    "user_name": "john_doe",
                                    "email": "%s",
                                    "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                }
                                """, email),
                        HttpStatus.OK
                ));
    }

    @GetMapping("/list/page={page}")
    @Async
    public CompletableFuture<ResponseEntity<String>> get_profiles(@PathVariable Integer page) {
        return CompletableFuture.completedFuture(
                new ResponseEntity<>(
                        """
                                [
                                    {
                                        "user_name": "john_doe",
                                        "email": "john.doe@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "alice_smith",
                                        "email": "alice.smith@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "bob_jones",
                                        "email": "bob.jones@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "emma_wilson",
                                        "email": "emma.wilson@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "charlie_brown",
                                        "email": "charlie.brown@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "olivia_miller",
                                        "email": "olivia.miller@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "david_wright",
                                        "email": "david.wright@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "sophia_hall",
                                        "email": "sophia.hall@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "michael_smith",
                                        "email": "michael.smith@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    },
                                    {
                                        "user_name": "samantha_davis",
                                        "email": "samantha.davis@example.com",
                                        "profile_image": "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"
                                    }
                                ]
                                """
                        , HttpStatus.OK
                ));
    }

    @PutMapping("/user/email={email}")
    @Async
    public CompletableFuture<ResponseEntity<String>> edit_profile(@PathVariable String email, @RequestBody String profile) {
        return CompletableFuture.completedFuture(
                new ResponseEntity<>(
                        "profile edited: " + profile,
                        HttpStatus.OK
                ));
    }

    @DeleteMapping("/user/email={email}")
    @Async
    public CompletableFuture<ResponseEntity<String>> delete_profile(@PathVariable String email) {
        return CompletableFuture.completedFuture(
                new ResponseEntity<>(
                        "profile deleted: " + email,
                        HttpStatus.OK
                ));
    }

}

//@RestController
//@RequestMapping(path = "/profile")
//public class ProfilesController {
//
//    @GetMapping("/user/email={email}")
//    @Async
//    public CompletableFuture<ResponseEntity<Profile>> get_user_profile(@PathVariable String email) {
//        return CompletableFuture.completedFuture(new ResponseEntity<>(
//                new Profile("john_doe", "john.doe@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                HttpStatus.OK));
//    }
//
//    @GetMapping("/list/page={page}")
//    @Async
//    public CompletableFuture<ResponseEntity<List<Profile>>> get_profiles(@PathVariable Integer page) {
//        List<Profile> profilesList = Arrays.asList(
//                new Profile("john_doe", "john.doe@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("alice_smith", "alice.smith@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("bob_jones", "bob.jones@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("emma_wilson", "emma.wilson@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("charlie_brown", "charlie.brown@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("olivia_miller", "olivia.miller@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("david_wright", "david.wright@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("sophia_hall", "sophia.hall@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("michael_smith", "michael.smith@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg"),
//                new Profile("samantha_davis", "samantha.davis@example.com", "https://samequizy.pl/wp-content/uploads/2021/12/29/images_e2bca6c20b784ca3.jpeg")
//        );
//
//        return CompletableFuture.completedFuture(new ResponseEntity<>(
//                profilesList
//                , HttpStatus.OK));
//    }
//
//    @PutMapping("/user/email={email}")
//    @Async
//    public CompletableFuture<ResponseEntity<Profile>> edit_profile(@RequestBody Profile profile) {
//        return CompletableFuture.completedFuture(new ResponseEntity<>(profile, HttpStatus.OK));
//    }
//}
