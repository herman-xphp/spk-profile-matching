import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { API } from '../App';
import { Button } from '../components/ui/button';
import { Label } from '../components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../components/ui/select';
import { toast } from 'sonner';
import { Calculator, TrendingUp } from 'lucide-react';

const Perhitungan = () => {
  const navigate = useNavigate();
  const [jabatanList, setJabatanList] = useState([]);
  const [selectedJabatan, setSelectedJabatan] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchJabatan();
  }, []);

  const fetchJabatan = async () => {
    try {
      const response = await axios.get(`${API}/jabatan`);
      setJabatanList(response.data);
    } catch (error) {
      toast.error('Gagal memuat data jabatan');
    }
  };

  const handleCalculate = async () => {
    if (!selectedJabatan) {
      toast.error('Pilih jabatan terlebih dahulu');
      return;
    }

    setLoading(true);
    try {
      await axios.post(`${API}/profile-matching/calculate`, { 
        jabatan_id: parseInt(selectedJabatan) 
      });
      toast.success('Perhitungan berhasil!');
      navigate(`/hasil-ranking/${selectedJabatan}`);
    } catch (error) {
      toast.error(error.response?.data?.error || error.response?.data?.detail || 'Perhitungan gagal');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">Perhitungan Profile Matching</h1>
        <p className="text-gray-500 mt-1">Jalankan perhitungan untuk mendapatkan ranking tenaga kerja</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Calculator className="w-5 h-5" />
              Pilih Jabatan
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label>Jabatan</Label>
              <Select value={selectedJabatan} onValueChange={setSelectedJabatan}>
                <SelectTrigger data-testid="select-jabatan-perhitungan">
                  <SelectValue placeholder="Pilih Jabatan" />
                </SelectTrigger>
                <SelectContent>
                  {jabatanList.map((j) => (
                    <SelectItem key={j.id} value={String(j.id)}>{j.nama}</SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>

            <Button
              onClick={handleCalculate}
              disabled={loading || !selectedJabatan}
              data-testid="button-calculate"
              className="w-full"
            >
              {loading ? 'Menghitung...' : 'Jalankan Perhitungan'}
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <TrendingUp className="w-5 h-5" />
              Tentang Metode
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4 text-sm text-gray-600">
              <div>
                <h4 className="font-semibold text-gray-900 mb-2">Profile Matching</h4>
                <p>
                  Metode yang membandingkan kompetensi individu dengan profil jabatan yang dibutuhkan,
                  sehingga diketahui perbedaan (GAP) antara keduanya.
                </p>
              </div>
              <div>
                <h4 className="font-semibold text-gray-900 mb-2">Langkah Perhitungan:</h4>
                <ol className="list-decimal list-inside space-y-1">
                  <li>Hitung GAP (Nilai Aktual - Target)</li>
                  <li>Konversi GAP ke Bobot Nilai</li>
                  <li>Hitung Core Factor (CF) & Secondary Factor (SF)</li>
                  <li>Hitung skor per aspek dengan bobot CF 60%, SF 40%</li>
                  <li>Hitung skor total berdasarkan persentase aspek</li>
                  <li>Tentukan ranking berdasarkan skor tertinggi</li>
                </ol>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default Perhitungan;
