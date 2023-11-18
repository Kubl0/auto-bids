package autobids.apigateway.profiles.models;

public class Profile {

    private String email;
    private String user_name;
    private String profile_image;

    public Profile(String email, String user_name, String profile_image) {
        this.email = email;
        this.user_name = user_name;
        this.profile_image = profile_image;
    }

    public String get_email() {
        return email;
    }

    public String get_user_name() {
        return user_name;
    }

    public String get_profile_image() {
        return profile_image;
    }

    @Override
    public String toString() {
        return "{"+email+","+user_name+","+profile_image+"}";
    }
}
