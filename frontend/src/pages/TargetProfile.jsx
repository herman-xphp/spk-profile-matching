import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { API, AuthContext } from '../App';
import { Button } from '../components/ui/button';
import { Input } from '../components/ui/input';
import { Label } from '../components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '../components/ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../components/ui/select';
import { toast } from 'sonner';
import { Plus, Edit, Trash2 } from 'lucide-react';

const TargetProfile = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [jabatanList, setJabatanList] = useState([]);
  const [kriteriaList, setKriteriaList] = useState([]);
  const [aspekList, setAspekList] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({ id: '', jabatan_id: '', kriteria_id: '', target_nilai: '' });
  const [filterJabatan, setFilterJabatan] = useState('');

  useEffect(() => {
    fetchData();
    fetchJabatan();
    fetchKriteria();
    fetchAspek();
  }, []);

  useEffect(() => {
    if (filterJabatan && filterJabatan !== 'all') {
      fetchDataByJabatan(filterJabatan);
    } else {
      fetchData();
    }
  }, [filterJabatan]);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/target-profiles`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data target profile');
    }
  };

  const fetchDataByJabatan = async (jabatanId) => {
    try {
      const response = await axios.get(`${API}/target-profiles?jabatan_id=${jabatanId}`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data target profile');
    }
  };

  const fetchJabatan = async () => {
    try {
      const response = await axios.get(`${API}/jabatan`);
      setJabatanList(response.data);
    } catch (error) {
      console.error('Error fetching jabatan');
    }
  };

  const fetchKriteria = async () => {
    try {
      const response = await axios.get(`${API}/kriteria`);
      setKriteriaList(response.data);
    } catch (error) {
      console.error('Error fetching kriteria');
    }
  };

  const fetchAspek = async () => {
    try {
      const response = await axios.get(`${API}/aspek`);
      setAspekList(response.data);
    } catch (error) {
      console.error('Error fetching aspek');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      jabatan_id: parseInt(formData.jabatan_id),
      kriteria_id: parseInt(formData.kriteria_id),
      target_nilai: parseFloat(formData.target_nilai)
    };
    try {
      if (editMode) {
        await axios.put(`${API}/target-profiles/${formData.id}`, payload);
        toast.success('Target profile berhasil diperbarui');
      } else {
        await axios.post(`${API}/target-profiles`, payload);
        toast.success('Target profile berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      if (filterJabatan) {
        fetchDataByJabatan(filterJabatan);
      } else {
        fetchData();
      }
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus target profile ini?')) return;
    try {
      await axios.delete(`${API}/target-profiles/${id}`);
      toast.success('Target profile berhasil dihapus');
      if (filterJabatan) {
        fetchDataByJabatan(filterJabatan);
      } else {
        fetchData();
      }
    } catch (error) {
      toast.error('Gagal menghapus target profile');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', jabatan_id: '', kriteria_id: '', target_nilai: '' });
    setEditMode(false);
  };

  const handleEdit = (item) => {
    setFormData({
      ...item,
      jabatan_id: String(item.jabatan_id),
      kriteria_id: String(item.kriteria_id)
    });
    setEditMode(true);
    setOpen(true);
  };

  const handleAdd = () => {
    resetForm();
    if (filterJabatan) {
      setFormData({ ...formData, jabatan_id: filterJabatan });
    }
    setOpen(true);
  };

  const getJabatanName = (id) => jabatanList.find(j => j.id === id)?.nama || '-';
  const getKriteriaInfo = (id) => {
    const kriteria = kriteriaList.find(k => k.id === id);
    if (!kriteria) return '-';
    const aspek = aspekList.find(a => a.id === kriteria.aspek_id);
    return `${kriteria.kode} - ${kriteria.nama} (${aspek?.nama || '-'})`;
  };

  const isAdmin = user?.role === 'admin';

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Target Profile</h1>
          <p className="text-gray-500 mt-1">Tentukan nilai target per kriteria untuk setiap jabatan</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-target-profile-button">
                <Plus size={16} className="mr-2" />
                Tambah Target
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Target Profile' : 'Tambah Target Profile'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label>Jabatan</Label>
                  <Select
                    value={formData.jabatan_id}
                    onValueChange={(value) => setFormData({ ...formData, jabatan_id: value })}
                    required
                  >
                    <SelectTrigger data-testid="tp-jabatan-select">
                      <SelectValue placeholder="Pilih Jabatan" />
                    </SelectTrigger>
                    <SelectContent>
                      {jabatanList.map((j) => (
                        <SelectItem key={j.id} value={String(j.id)}>{j.nama}</SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label>Kriteria</Label>
                  <Select
                    value={formData.kriteria_id}
                    onValueChange={(value) => setFormData({ ...formData, kriteria_id: value })}
                    required
                  >
                    <SelectTrigger data-testid="tp-kriteria-select">
                      <SelectValue placeholder="Pilih Kriteria" />
                    </SelectTrigger>
                    <SelectContent>
                      {kriteriaList.map((k) => (
                        <SelectItem key={k.id} value={String(k.id)}>
                          {k.kode} - {k.nama}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="target_nilai">Nilai Target</Label>
                  <Input
                    id="target_nilai"
                    data-testid="tp-nilai-input"
                    type="number"
                    step="0.1"
                    value={formData.target_nilai}
                    onChange={(e) => setFormData({ ...formData, target_nilai: e.target.value })}
                    required
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="tp-submit-button" className="flex-1">
                    {editMode ? 'Update' : 'Simpan'}
                  </Button>
                  <Button type="button" variant="outline" onClick={() => setOpen(false)} className="flex-1">
                    Batal
                  </Button>
                </div>
              </form>
            </DialogContent>
          </Dialog>
        )}
      </div>

      <div className="mb-4">
        <Label>Filter berdasarkan Jabatan</Label>
        <Select value={filterJabatan} onValueChange={setFilterJabatan}>
          <SelectTrigger className="w-full max-w-md" data-testid="filter-jabatan-select">
            <SelectValue placeholder="Semua Jabatan" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Semua Jabatan</SelectItem>
            {jabatanList.map((j) => (
              <SelectItem key={j.id} value={String(j.id)}>{j.nama}</SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Daftar Target Profile</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data target profile</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Jabatan</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Kriteria</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Target Nilai</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`tp-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4 font-medium">{getJabatanName(item.jabatan_id)}</td>
                      <td className="py-3 px-4">{getKriteriaInfo(item.kriteria_id)}</td>
                      <td className="py-3 px-4 text-gray-600">{item.target_nilai}</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button size="sm" variant="outline" onClick={() => handleEdit(item)} data-testid={`edit-tp-${index}`}>
                              <Edit size={14} />
                            </Button>
                            <Button size="sm" variant="destructive" onClick={() => handleDelete(item.id)} data-testid={`delete-tp-${index}`}>
                              <Trash2 size={14} />
                            </Button>
                          </div>
                        </td>
                      )}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default TargetProfile;
