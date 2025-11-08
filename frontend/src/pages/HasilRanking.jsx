import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate, Link } from 'react-router-dom';
import { API } from '../App';
import { Button } from '../components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { toast } from 'sonner';
import { Trophy, Eye, Download, ArrowLeft } from 'lucide-react';

const HasilRanking = () => {
  const { jabatanId } = useParams();
  const navigate = useNavigate();
  const [results, setResults] = useState([]);
  const [jabatan, setJabatan] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchResults();
    fetchJabatan();
  }, [jabatanId]);

  const fetchResults = async () => {
    try {
      const response = await axios.get(`${API}/profile-matching/results?jabatan_id=${jabatanId}`);
      setResults(response.data);
    } catch (error) {
      toast.error('Gagal memuat hasil ranking');
    } finally {
      setLoading(false);
    }
  };

  const fetchJabatan = async () => {
    try {
      const response = await axios.get(`${API}/jabatan`);
      const jabatanIdNum = parseInt(jabatanId);
      const j = response.data.find(j => j.id === jabatanIdNum);
      setJabatan(j);
    } catch (error) {
      console.error('Error fetching jabatan', error);
    }
  };

  const handleExport = () => {
    const csv = [
      ['Rank', 'NIK', 'Nama', 'Skor Total'],
      ...results.map(r => [
        r.rank,
        r.tenaga_kerja?.nik || '-',
        r.tenaga_kerja?.nama || '-',
        r.score_total
      ])
    ].map(row => row.join(',')).join('\n');

    const blob = new Blob([csv], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `hasil_ranking_${jabatan?.nama || 'jabatan'}.csv`;
    a.click();
    toast.success('Data berhasil diexport');
  };

  const getRankColor = (rank) => {
    if (rank === 1) return 'bg-yellow-100 text-yellow-800';
    if (rank === 2) return 'bg-gray-100 text-gray-800';
    if (rank === 3) return 'bg-orange-100 text-orange-800';
    return 'bg-blue-50 text-blue-700';
  };

  if (loading) {
    return <div className="text-center py-8">Memuat...</div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <Button
            variant="ghost"
            onClick={() => navigate('/perhitungan')}
            className="mb-2"
            data-testid="back-button"
          >
            <ArrowLeft size={16} className="mr-2" />
            Kembali
          </Button>
          <h1 className="text-2xl font-bold text-gray-900">Hasil Ranking</h1>
          <p className="text-gray-500 mt-1">Jabatan: {jabatan?.nama || '-'}</p>
        </div>
        <Button onClick={handleExport} data-testid="export-button">
          <Download size={16} className="mr-2" />
          Export CSV
        </Button>
      </div>

      {results.length === 0 ? (
        <Card>
          <CardContent className="py-12">
            <p className="text-center text-gray-500">Belum ada hasil perhitungan</p>
          </CardContent>
        </Card>
      ) : (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Trophy className="w-5 h-5 text-yellow-500" />
              Ranking Tenaga Kerja
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Rank</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">NIK</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nama</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Skor Total</th>
                    <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  {results.map((result, index) => (
                    <tr key={result.id} className="border-b hover:bg-gray-50" data-testid={`result-row-${index}`}>
                      <td className="py-3 px-4">
                        <span className={`px-3 py-1 rounded-full text-sm font-semibold ${getRankColor(result.rank)}`}>
                          #{result.rank}
                        </span>
                      </td>
                      <td className="py-3 px-4 font-medium">{result.tenaga_kerja?.nik || '-'}</td>
                      <td className="py-3 px-4">{result.tenaga_kerja?.nama || '-'}</td>
                      <td className="py-3 px-4">
                        <span className="text-lg font-bold text-indigo-600">
                          {result.score_total}
                        </span>
                      </td>
                      <td className="py-3 px-4">
                        <div className="flex justify-center">
                          <Link to={`/detail-hasil/${result.id}`}>
                            <Button size="sm" variant="outline" data-testid={`view-detail-${index}`}>
                              <Eye size={14} className="mr-1" />
                              Detail
                            </Button>
                          </Link>
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
};

export default HasilRanking;
