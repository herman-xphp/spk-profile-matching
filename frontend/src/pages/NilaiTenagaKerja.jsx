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

const NilaiTenagaKerja = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [tenagaKerjaList, setTenagaKerjaList] = useState([]);
  const [kriteriaList, setKriteriaList] = useState([]);
  const [aspekList, setAspekList] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({ id: '', tenaga_kerja_id: '', kriteria_id: '', nilai: '' });
  const [filterTenagaKerja, setFilterTenagaKerja] = useState('');

  useEffect(() => {
    fetchData();
    fetchTenagaKerja();
    fetchKriteria();
    fetchAspek();
  }, []);

  useEffect(() => {
    if (filterTenagaKerja && filterTenagaKerja !== 'all') {
      fetchDataByTenagaKerja(filterTenagaKerja);
    } else {
      fetchData();
    }
  }, [filterTenagaKerja]);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/nilai-tenaga-kerja`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data nilai');
    }
  };

  const fetchDataByTenagaKerja = async (tkId) => {
    try {
      const response = await axios.get(`${API}/nilai-tenaga-kerja?tenaga_kerja_id=${tkId}`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data nilai');
    }
  };

  const fetchTenagaKerja = async () => {
    try {
      const response = await axios.get(`${API}/tenaga-kerja`);
      setTenagaKerjaList(response.data);
    } catch (error) {
      console.error('Error fetching tenaga kerja', error);
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
      tenaga_kerja_id: parseInt(formData.tenaga_kerja_id),
      kriteria_id: parseInt(formData.kriteria_id),
      nilai: parseFloat(formData.nilai)
    };
    try {
      if (editMode) {
        await axios.put(`${API}/nilai-tenaga-kerja/${formData.id}`, payload);
        toast.success('Nilai berhasil diperbarui');
      } else {
        await axios.post(`${API}/nilai-tenaga-kerja`, payload);
        toast.success('Nilai berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      if (filterTenagaKerja && filterTenagaKerja !== 'all') {
        fetchDataByTenagaKerja(filterTenagaKerja);
      } else {
        fetchData();
      }
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus nilai ini?')) return;
    try {
      await axios.delete(`${API}/nilai-tenaga-kerja/${id}`);
      toast.success('Nilai berhasil dihapus');
      if (filterTenagaKerja && filterTenagaKerja !== 'all') {
        fetchDataByTenagaKerja(filterTenagaKerja);
      } else {
        fetchData();
      }
    } catch (error) {
      toast.error('Gagal menghapus nilai');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', tenaga_kerja_id: '', kriteria_id: '', nilai: '' });
    setEditMode(false);
  };

  const handleEdit = (item) => {
    setFormData({
      ...item,
      tenaga_kerja_id: String(item.tenaga_kerja_id),
      kriteria_id: String(item.kriteria_id)
    });
    setEditMode(true);
    setOpen(true);
  };

  const handleAdd = () => {
    resetForm();
    if (filterTenagaKerja) {
      setFormData({ ...formData, tenaga_kerja_id: filterTenagaKerja });
    }
    setOpen(true);
  };

  const getTenagaKerjaName = (id) => {
    const tk = tenagaKerjaList.find(t => t.id === id);
    return tk ? `${tk.nik} - ${tk.nama}` : '-';
  };

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
          <h1 className="text-2xl font-bold text-gray-900">Nilai Tenaga Kerja</h1>
          <p className="text-gray-500 mt-1">Input nilai aktual tenaga kerja per kriteria</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-nilai-button">
                <Plus size={16} className="mr-2" />
                Tambah Nilai
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Nilai' : 'Tambah Nilai'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label>Tenaga Kerja</Label>
                  <Select
                    value={formData.tenaga_kerja_id}
                    onValueChange={(value) => setFormData({ ...formData, tenaga_kerja_id: value })}
                    required
                  >
                    <SelectTrigger data-testid="nilai-tk-select">
                      <SelectValue placeholder="Pilih Tenaga Kerja" />
                    </SelectTrigger>
                    <SelectContent>
                      {tenagaKerjaList.map((tk) => (
                        <SelectItem key={tk.id} value={String(tk.id)}>{tk.nik} - {tk.nama}</SelectItem>
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
                    <SelectTrigger data-testid="nilai-kriteria-select">
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
                  <Label htmlFor="nilai">Nilai</Label>
                  <Input
                    id="nilai"
                    data-testid="nilai-input"
                    type="number"
                    step="0.1"
                    value={formData.nilai}
                    onChange={(e) => setFormData({ ...formData, nilai: e.target.value })}
                    required
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="nilai-submit-button" className="flex-1">
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
        <Label>Filter berdasarkan Tenaga Kerja</Label>
        <Select value={filterTenagaKerja} onValueChange={setFilterTenagaKerja}>
          <SelectTrigger className="w-full max-w-md" data-testid="filter-tk-select">
            <SelectValue placeholder="Semua Tenaga Kerja" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Semua Tenaga Kerja</SelectItem>
            {tenagaKerjaList.map((tk) => (
              <SelectItem key={tk.id} value={String(tk.id)}>{tk.nik} - {tk.nama}</SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Daftar Nilai Tenaga Kerja</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data nilai</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Tenaga Kerja</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Kriteria</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nilai</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`nilai-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4 font-medium">{getTenagaKerjaName(item.tenaga_kerja_id)}</td>
                      <td className="py-3 px-4">{getKriteriaInfo(item.kriteria_id)}</td>
                      <td className="py-3 px-4 text-gray-600">{item.nilai}</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button size="sm" variant="outline" onClick={() => handleEdit(item)} data-testid={`edit-nilai-${index}`}>
                              <Edit size={14} />
                            </Button>
                            <Button size="sm" variant="destructive" onClick={() => handleDelete(item.id)} data-testid={`delete-nilai-${index}`}>
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

export default NilaiTenagaKerja;
