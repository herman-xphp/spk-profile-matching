import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import { API } from '../App';
import { Button } from '../components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { toast } from 'sonner';
import { ArrowLeft, User, Briefcase } from 'lucide-react';

const DetailHasil = () => {
  const { resultId } = useParams();
  const navigate = useNavigate();
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchDetail();
  }, [resultId]);

  const fetchDetail = async () => {
    try {
      const response = await axios.get(`${API}/profile-matching/results/${resultId}`);
      setResult(response.data);
    } catch (error) {
      toast.error('Gagal memuat detail hasil');
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div className="text-center py-8">Memuat...</div>;
  }

  if (!result) {
    return <div className="text-center py-8">Data tidak ditemukan</div>;
  }

  const details = result.details || {};
  const aspekData = details.aspek || {};

  return (
    <div>
      <Button
        variant="ghost"
        onClick={() => navigate(-1)}
        className="mb-4"
        data-testid="back-button-detail"
      >
        <ArrowLeft size={16} className="mr-2" />
        Kembali
      </Button>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-lg">
              <User className="w-5 h-5" />
              Tenaga Kerja
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2 text-sm">
              <div>
                <span className="text-gray-500">NIK:</span>
                <p className="font-medium">{result.tenaga_kerja?.nik || '-'}</p>
              </div>
              <div>
                <span className="text-gray-500">Nama:</span>
                <p className="font-medium">{result.tenaga_kerja?.nama || '-'}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-lg">
              <Briefcase className="w-5 h-5" />
              Jabatan
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2 text-sm">
              <div>
                <span className="text-gray-500">Nama Jabatan:</span>
                <p className="font-medium">{result.jabatan?.nama || '-'}</p>
              </div>
              <div>
                <span className="text-gray-500">Deskripsi:</span>
                <p className="font-medium">{result.jabatan?.deskripsi || '-'}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="text-lg">Hasil Akhir</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2 text-sm">
              <div>
                <span className="text-gray-500">Ranking:</span>
                <p className="text-2xl font-bold text-indigo-600">#{result.rank}</p>
              </div>
              <div>
                <span className="text-gray-500">Skor Total:</span>
                <p className="text-2xl font-bold text-indigo-600">{result.score_total}</p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Detail Perhitungan per Aspek</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-6">
            {Object.entries(aspekData).map(([aspekNama, aspekDetail]) => (
              <div key={aspekNama} className="border rounded-lg p-4">
                <div className="flex justify-between items-center mb-4">
                  <h3 className="text-lg font-semibold text-gray-900">{aspekNama}</h3>
                  <div className="text-right">
                    <p className="text-sm text-gray-500">Persentase: {aspekDetail.persentase}%</p>
                    <p className="text-sm text-gray-500">Skor: {aspekDetail.score?.toFixed(2)}</p>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-4 mb-4">
                  <div className="bg-blue-50 p-3 rounded">
                    <p className="text-xs text-gray-600">Core Factor (CF)</p>
                    <p className="text-xl font-bold text-blue-600">{aspekDetail.cf?.toFixed(2)}</p>
                  </div>
                  <div className="bg-green-50 p-3 rounded">
                    <p className="text-xs text-gray-600">Secondary Factor (SF)</p>
                    <p className="text-xl font-bold text-green-600">{aspekDetail.sf?.toFixed(2)}</p>
                  </div>
                </div>

                <div className="overflow-x-auto">
                  <table className="w-full text-sm">
                    <thead>
                      <tr className="border-b">
                        <th className="text-left py-2 px-2 font-semibold text-gray-700">Kode</th>
                        <th className="text-left py-2 px-2 font-semibold text-gray-700">Kriteria</th>
                        <th className="text-center py-2 px-2 font-semibold text-gray-700">Target</th>
                        <th className="text-center py-2 px-2 font-semibold text-gray-700">Aktual</th>
                        <th className="text-center py-2 px-2 font-semibold text-gray-700">GAP</th>
                        <th className="text-center py-2 px-2 font-semibold text-gray-700">Bobot</th>
                        <th className="text-center py-2 px-2 font-semibold text-gray-700">Tipe</th>
                      </tr>
                    </thead>
                    <tbody>
                      {aspekDetail.kriteria?.map((k, idx) => (
                        <tr key={idx} className="border-b">
                          <td className="py-2 px-2 font-medium">{k.kode}</td>
                          <td className="py-2 px-2">{k.nama}</td>
                          <td className="py-2 px-2 text-center">{k.target}</td>
                          <td className="py-2 px-2 text-center">{k.actual}</td>
                          <td className="py-2 px-2 text-center">
                            <span className={`px-2 py-1 rounded ${k.gap === 0 ? 'bg-green-100 text-green-700' : k.gap > 0 ? 'bg-blue-100 text-blue-700' : 'bg-red-100 text-red-700'}`}>
                              {k.gap}
                            </span>
                          </td>
                          <td className="py-2 px-2 text-center font-semibold">{k.bobot_nilai}</td>
                          <td className="py-2 px-2 text-center">
                            <span className={`px-2 py-1 text-xs rounded-full ${k.is_core ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700'}`}>
                              {k.is_core ? 'Core' : 'Secondary'}
                            </span>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default DetailHasil;
