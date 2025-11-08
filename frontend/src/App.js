import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import axios from 'axios';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Jabatan from './pages/Jabatan';
import Aspek from './pages/Aspek';
import Kriteria from './pages/Kriteria';
import TenagaKerja from './pages/TenagaKerja';
import TargetProfile from './pages/TargetProfile';
import NilaiTenagaKerja from './pages/NilaiTenagaKerja';
import Perhitungan from './pages/Perhitungan';
import HasilRanking from './pages/HasilRanking';
import DetailHasil from './pages/DetailHasil';
import AdminLayout from './components/AdminLayout';
import { Toaster } from './components/ui/sonner';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;
export const API = `${BACKEND_URL}/api`;

export const AuthContext = React.createContext();

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      checkAuth();
    } else {
      setLoading(false);
    }
  }, []);

  const checkAuth = async () => {
    try {
      const response = await axios.get(`${API}/auth/me`);
      setUser(response.data);
    } catch (error) {
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];
    } finally {
      setLoading(false);
    }
  };

  const login = (token, userData) => {
    localStorage.setItem('token', token);
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    setUser(userData);
  };

  const logout = () => {
    localStorage.removeItem('token');
    delete axios.defaults.headers.common['Authorization'];
    setUser(null);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-lg">Memuat...</div>
      </div>
    );
  }

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={!user ? <Login /> : <Navigate to="/" />} />
          <Route path="/register" element={!user ? <Register /> : <Navigate to="/" />} />
          
          <Route element={user ? <AdminLayout /> : <Navigate to="/login" />}>
            <Route path="/" element={<Dashboard />} />
            <Route path="/jabatan" element={<Jabatan />} />
            <Route path="/aspek" element={<Aspek />} />
            <Route path="/kriteria" element={<Kriteria />} />
            <Route path="/tenaga-kerja" element={<TenagaKerja />} />
            <Route path="/target-profile" element={<TargetProfile />} />
            <Route path="/nilai-tenaga-kerja" element={<NilaiTenagaKerja />} />
            <Route path="/perhitungan" element={<Perhitungan />} />
            <Route path="/hasil-ranking/:jabatanId" element={<HasilRanking />} />
            <Route path="/detail-hasil/:resultId" element={<DetailHasil />} />
          </Route>
        </Routes>
      </BrowserRouter>
      <Toaster position="top-right" />
    </AuthContext.Provider>
  );
}

export default App;
