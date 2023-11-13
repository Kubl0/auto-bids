import './App.css';
import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer';
import MainPage from './components/MainPage';
import {BrowserRouter as Router} from 'react-router-dom';
import { Routes, Route } from 'react-router-dom';
import SearchPage from './components/SearchPage';
import Account from './components/Account';
import NewListing from './components/NewListing';
import LoginPage from './components/LoginPage';
import RegisterPage from './components/RegisterPage';
import OfferPage from './components/OfferPage';
import PrivacyPolicy from './components/PrivacyPolicy';
import AboutUs from './components/AboutUs';
import ContactUs from './components/ContactUs';
import TermsOfService from './components/TermsOfService';
import Help from './components/Help';
import SellerPage from './components/SellerPage';

function App() {
  return (
    <div className="App">
      <Router>
        <Header />
           <Routes>
              <Route path="/" element={<MainPage />} />
              <Route path="/search" element={<SearchPage />} />
              <Route path="/account" element={<Account />} />
              <Route path="/new-listing" element={<NewListing />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/register" element={<RegisterPage />} />
              <Route path="/offer" element={<OfferPage />} />
              <Route path="/privacy-policy" element={<PrivacyPolicy />} />
              <Route path="/about-us" element={<AboutUs />} />
              <Route path="/contact-us" element={<ContactUs />} />
              <Route path="/terms-of-service" element={<TermsOfService />} />
              <Route path="/help" element={<Help />} />
              <Route path="/seller" element={<SellerPage />} />
           </Routes>
        <Footer />
      </Router>
    </div>
  );
}

export default App;
