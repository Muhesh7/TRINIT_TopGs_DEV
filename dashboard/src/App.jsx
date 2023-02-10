import { Routes, Route } from "react-router-dom"
import { LoginPage } from "./pages/Login";
import { SignupPage } from "./pages/Signup";
import { HomePage } from "./pages/Home";
import { ProfilePage } from "./pages/Profile";
import { SettingsPage } from "./pages/Settings";
import { ProtectedLayout } from "./components/ProtectedLayout";
import { HomeLayout } from "./components/HomeLayout";

export default function App() {
  return (
    <Routes>
      <Route element={<HomeLayout />}>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/signup" element={<SignupPage />} />
      </Route>

      <Route path="/dashboard" element={<ProtectedLayout />}>
        <Route path="profile" element={<ProfilePage />} />
        <Route path="settings" element={<SettingsPage />} />
      </Route>
    </Routes>
  );
}
