import React from "react";
import OfferElement from "./OfferElement";
import Chat from "./Chat";
import { useEffect, useState } from "react";
import * as Yup from "yup";
import { Link } from "react-router-dom";
import EditOffer from "./EditOffer";

interface IOffer {
  mileage: number;
  photos: string[];
  description: string;
  id: string;
  image: string;
  title: string;
  price: number;
  year: number;
};


interface IProfile {
    email: string;
    user_name: string;
    profile_picture: string;
}

const AccountSchema = Yup.object().shape({
  user_name: Yup.string()
    .min(2, "Too Short!")
    .max(50, "Too Long!")
    .required("Required"),
  profile_picture: Yup.string()
    .url("Invalid URL")
    .required("Required"),
});

export default function Account({ setIsLoggedIn }: {setIsLoggedIn: (value: boolean) => void;}) {
    const [isEditing, setIsEditing] = useState(false);
    const [editedProfile, setEditedProfile] = useState<IProfile | null>(null);
    const [pageNumber, setPageNumber] = useState(1);
    const [editingOfferId, setEditingOfferId] = useState<string | null>(null)

    const handleEditProfile = () => {
      setIsEditing(true);
      setEditedProfile(profileData ? { ...profileData } : null);
    };

    const handleCancelEdit = () => {
      setIsEditing(false);
      setEditedProfile(null);
    };

    const handleDeleteAllOffers = () => {
      const confirmed = window.confirm("Are you sure you want to delete all your offers?");
      if (confirmed) {
        try{
          fetch(`${process.env.REACT_APP_PROFILE_DELETE_ENDPOINT}`, {
            method: "DELETE",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              "Access-Control-Allow-Credentials": "true",
            },
          });
        } catch (error) {
          console.error("Error deleting all offers:", error);
        }
      }
    };
    

    const handleSaveProfile = () => {
      AccountSchema.validate(editedProfile)
        .then(() => {
          setIsEditing(false);
          setProfileData(editedProfile ? { ...editedProfile } : null);
          const dataToSend = {
            profile_image: editedProfile?.profile_picture,
            user_name: editedProfile?.user_name,
          };
          fetch(`${process.env.REACT_APP_PROFILE_EDIT_ENDPOINT}`, {
            method: "PUT",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              "Access-Control-Allow-Credentials": "true",
            },
            body: JSON.stringify(dataToSend),
          });
        })
        .catch((error) => {
          console.error("Error validating profile data:", error);
          if (error instanceof Yup.ValidationError) {
            if (error.path === "user_name") {
              alert("Name must be between 2 and 50 characters long");
            }
            if (error.path === "profile_picture") {
              alert("Profile picture must be a valid URL");
            }
          }
        });
    }

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      const { name, value } = e.target;
      setEditedProfile(editedProfile ? { ...editedProfile, [name]: value } : null);
    };

    const [offerData, setOfferData] = useState<IOffer []| null>(null);
    const [profileData, setProfileData] = useState<IProfile | null>(null);

    const handleNextPage = () => {
      setPageNumber(pageNumber + 1);
    }
    const handlePreviousPage = () => {
      if (pageNumber > 1) {
        setPageNumber(pageNumber - 1);
      }
    }

      const fetchData = async () => {
        try {
          const response = await fetch(`${process.env.REACT_APP_PROFILE_LOGIN_ENDPOINT}`, {
            method: "GET",
                credentials: "include",
                headers: {
                    "Access-Control-Allow-Origin": "*",
                    "Access-Control-Allow-Credentials": "true",
                },
          });
          if (response.status === 201) {
            document.cookie = "isLoggedIn=true";
            window.location.reload();
          }
          if(response.status === 200){
            setIsLoggedIn(true);
            document.cookie = "isLoggedIn=true";
            const profileData = await response.json();
            setProfileData({
              email: profileData.data.data.email,
              user_name: profileData.data.data.user_name,
              profile_picture: profileData.data.data.profile_image,
            });
          }
        } catch (error) {
          console.error("Error loading profile data:", error);
        }
      };

      const fetchUsersOffers = async () => {
        try {
          const response = await fetch(`${process.env.REACT_APP_PROFILE_CARS_ENDPOINT}${pageNumber}`, {
            method: "GET",
            credentials: "include",
            headers: {
              "Access-Control-Allow-Origin": "*",
              "Access-Control-Allow-Credentials": "true",
            },
          });
          if (response.ok) {
            const offers = await response.json();
            const offerData: IOffer[] = [];
            offers.data.data.forEach((offer: any) => {
              offerData.push({
                id: offer.id,
                image: offer.car.photos.length > 0 ? offer.car.photos[0] : "",
                title: offer.car.title,
                price: offer.car.price,
                year: offer.car.year,
                description: offer.car.description,
                mileage: offer.car.mileage,
                photos: offer.car.photos,
              });
            });
            setOfferData(offerData);
          } else {
            console.log("Error fetching offers");
          }
        } catch (error) {
          console.error("Error fetching offers:", error);
        }
      }

      useEffect(() => {
        fetchData();
        // fetchUsersOffers();
      }, []);

      useEffect(() => {
        fetchUsersOffers();
      }, [pageNumber]);

      const handleDeleteProfile = async () => {
        const confirmed = window.confirm("Are you sure you want to delete your account?");

        if (confirmed) {
          try {
            const response = await fetch(`${process.env.REACT_APP_CARS_DELETE_ALL_USER_CARS_ENDPOINT}`, {
              method: "DELETE",
              credentials: "include",
              headers: {
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Credentials": "true",
              },
            });
      
            if (response.ok) {
              console.log("All offers deleted");
            } else {
              console.log("Error deleting all offers");
            }
          } catch (error) {
            console.error("Error deleting all offers:", error);
          }
          try {
            const response = await fetch(`${process.env.REACT_APP_PROFILE_DELETE_ENDPOINT}`, {
              method: "DELETE",
              credentials: "include",
              headers: {
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Credentials": "true",
              },
            });
      
            const response2 = await fetch(`${process.env.REACT_APP_LOGOUT_ENDPOINT}`, {
              method: "POST",
              credentials: "include",
              headers: {
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Credentials": "true",
              },
            });
      
            if (response.ok && response2.ok) {
              setIsLoggedIn(false);
              document.cookie = "isLoggedIn=false";
              window.location.href = "/";
            } else {
              console.log("Error deleting profile");
            }
          } catch (error) {
            console.error("Error deleting profile:", error);
          }
        }
      };

      const handleLogout = async () => {
        try{
          const response = await fetch(`${process.env.REACT_APP_LOGOUT_ENDPOINT}`, {
            method: "POST",
            credentials: "include",
            headers: {
              "Access-Control-Allow-Origin": "*",
              "Access-Control-Allow-Credentials": "true",
            },
          });
          if (response.ok){
            setIsLoggedIn(false);
            document.cookie = "isLoggedIn=false";
            window.location.href = "/";
          }
          else{
            console.log("Error logging out");
          }
        } catch (error) {
          console.error("Error logging out:", error);
        }
      };

      const handleDeleteOffer = async (id: string) => {
        const confirmed = window.confirm("Are you sure you want to delete this offer?");
        if (confirmed) {
          try {
            const response = await fetch(`${process.env.REACT_APP_CARS_DELETE_ENDPOINT}`, {
              method: "DELETE",
              credentials: "include",
              headers: {
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Credentials": "true",
              },
              body: JSON.stringify({ id }),
            });
            if (response.ok) {
              fetchUsersOffers();
            } else {
              console.log("Error deleting offer");
            }
          } catch (error) {
            console.error("Error deleting offer:", error);
          }
        }
      }

      const handleEditOffer = (id: string) => {
        setEditingOfferId(id);
      }

      return (
        <div className="account p-4">
          <div className="account-header flex justify-between items-center">
            <div className="account-header-profile flex items-center">
              <img src={profileData?.profile_picture} alt="profile" className="w-12 h-12 rounded-full mr-4" />
              <div className="account-header-profile-info">
                <h2 className="text-xl font-bold">{profileData?.user_name}</h2>
                <p>{profileData?.email}</p>
              </div>
            </div>
            <div className="account-header-buttons flex">
              <div className="account-header-buttons-element mr-4">
                <button onClick={handleEditProfile} className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition duration-300">
                  Edit Profile
                </button>
              </div>
              <div className="account-header-buttons-element mr-4">
                <button onClick={handleDeleteProfile} className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition duration-300">
                  Delete
                </button>
              </div>
              <div className="account-header-buttons-element">
                <button onClick={handleLogout} className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600 transition duration-300">
                  Logout
                </button>
              </div>
            </div>
          </div>
          {isEditing && editedProfile && (
            <div className="account-edit-profile mt-4 p-4 border border-gray-300 rounded">
              <h2 className="text-xl font-bold mb-4">Edit Profile</h2>
              <label className="block mb-2">
                Name:
                <input
                  type="text"
                  name="user_name"
                  value={editedProfile.user_name}
                  onChange={handleInputChange}
                  className="border border-gray-300 p-2 rounded w-full"
                />
              </label>
              <label className="block mb-4">
                Profile picture:
                <input
                  type="text"
                  name="profile_picture"
                  value={editedProfile.profile_picture}
                  onChange={handleInputChange}
                  className="border border-gray-300 p-2 rounded w-full"
                />
              </label>
              <div className="flex justify-end">
                <button onClick={handleCancelEdit} className="bg-gray-500 text-white px-4 py-2 rounded mr-2 hover:bg-gray-600 transition duration-300">
                  Cancel
                </button>
                <button onClick={handleSaveProfile} className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition duration-300">
                  Save
                </button>
              </div>
            </div>
          )}
          <div className="account-offers mt-4">
            <h2 className="text-xl font-bold mb-4">Your offers</h2>
            <button onClick={handleDeleteAllOffers} className="bg-red-500 text-white px-4 py-2 rounded mb-4 hover:bg-red-600 transition duration-300">
              Delete all offers
            </button>
            <div className="account-offers-elements mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {offerData &&
              offerData.map((offer) => (
                <div key={offer.id} className="border p-4 rounded">
                  {editingOfferId === offer.id ? (
                    <EditOffer
                      id={offer.id}
                      title={offer.title}
                      description={offer.description}
                      price={offer.price}
                      mileage={offer.mileage}
                      photos={offer.photos}
                    />
                  ) : (
                    <Link to={`/cars/offer/${offer.id}`}>
                      <OfferElement
                        image={offer.image}
                        title={offer.title}
                        price={offer.price}
                        year={offer.year}
                      />
                    </Link>
                  )}
                  <button
                    onClick={() => handleDeleteOffer(offer.id)}
                    className="mt-2 bg-red-500 text-white px-4 py-2 mr-2 rounded hover:bg-red-600 transition duration-300"
                  >
                    Delete
                  </button>
                  {editingOfferId === offer.id ? ( 
                    <button
                      onClick={() => handleEditOffer('')}
                      className="mt-2 bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600 transition duration-300"
                    >
                      Cancel
                    </button>
                  ) : (
                    <button
                      onClick={() => handleEditOffer(offer.id)}
                      className="mt-2 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition duration-300"
                    >
                      Edit
                    </button>
                  )}
                </div>
              ))}
          </div>
            <div className="mt-4 flex justify-between">
              <button onClick={handlePreviousPage} className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600 transition duration-300">
                Previous
              </button>
              <button onClick={handleNextPage} className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600 transition duration-300">
                Next
              </button>
            </div>
          </div>
          <div className="account-saved-offers mt-4">
            <h2 className="text-xl font-bold mb-4">Saved offers</h2>
            <div className="account-saved-offers-elements grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {offerData &&
                offerData.map((offer) => (
                  <div key={offer.id} className="border p-4 rounded">
                    <OfferElement key={offer.id} image={offer.image} title={offer.title} price={offer.price} year={offer.year} />
                  </div>
                ))}
            </div>
          </div>
          <div className="account-chat mt-4">
            <h2 className="text-xl font-bold mb-4">Chat</h2>
            <div className="account-chat-elements">
              <Chat />
            </div>
          </div>
        </div>
      );
      
}