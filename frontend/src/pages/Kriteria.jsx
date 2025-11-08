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

const Kriteria = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [aspekList, setAspekList] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({
    id: '', aspek_id: '', kode: '', nama: '', is_core: true, bobot: 1.0
  });

  useEffect(() => {
    fetchData();
    fetchAspek();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/kriteria`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data kriteria');
    }
  };

  const fetchAspek = async () => {
    try {
      const response = await axios.get(`${API}/aspek`);
      setAspekList(response.data);
    } catch (error) {
      toast.error('Gagal memuat data aspek');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      aspek_id: formData.aspek_id,
      kode: formData.kode,
      nama: formData.nama,
      is_core: formData.is_core,
      bobot: parseFloat(formData.bobot)
    };
    try {
      if (editMode) {
        await axios.put(`${API}/kriteria/${formData.id}`, payload);
        toast.success('Kriteria berhasil diperbarui');
      } else {
        await axios.post(`${API}/kriteria`, payload);
        toast.success('Kriteria berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      fetchData();
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus kriteria ini?')) return;
    try {
      await axios.delete(`${API}/kriteria/${id}`);
      toast.success('Kriteria berhasil dihapus');
      fetchData();
    } catch (error) {
      toast.error('Gagal menghapus kriteria');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', aspek_id: '', kode: '', nama: '', is_core: true, bobot: 1.0 });
    setEditMode(false);
  };

  const handleEdit = (item) => {
    setFormData(item);
    setEditMode(true);
    setOpen(true);
  };

  const handleAdd = () => {
    resetForm();
    setOpen(true);
  };

  const getAspekName = (aspekId) => {
    const aspek = aspekList.find(a => a.id === aspekId);
    return aspek ? aspek.nama : '-';
  };

  const isAdmin = user?.role === 'admin';

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Data Kriteria</h1>
          <p className="text-gray-500 mt-1">Kelola kriteria penilaian per aspek</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-kriteria-button">
                <Plus size={16} className="mr-2" />
                Tambah Kriteria
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Kriteria' : 'Tambah Kriteria'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="aspek">Aspek</Label>
                  <Select
                    value={formData.aspek_id}
                    onValueChange={(value) => setFormData({ ...formData, aspek_id: value })}
                    required
                  >
                    <SelectTrigger data-testid="kriteria-aspek-select">
                      <SelectValue placeholder="Pilih Aspek" />
                    </SelectTrigger>
                    <SelectContent>
                      {aspekList.map((aspek) => (
                        <SelectItem key={aspek.id} value={aspek.id}>{aspek.nama}</SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="kode">Kode</Label>
                  <Input
                    id="kode"
                    data-testid="kriteria-kode-input"
                    value={formData.kode}
                    onChange={(e) => setFormData({ ...formData, kode: e.target.value })}
                    placeholder="K1, S1, P1, dll"
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="nama">Nama Kriteria</Label>
                  <Input
                    id="nama"
                    data-testid="kriteria-nama-input"
                    value={formData.nama}
                    onChange={(e) => setFormData({ ...formData, nama: e.target.value })}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="is_core">Tipe Faktor</Label>
                  <Select
                    value={formData.is_core.toString()}
                    onValueChange={(value) => setFormData({ ...formData, is_core: value === 'true' })}
                  >
                    <SelectTrigger data-testid="kriteria-is-core-select">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="true">Core Factor</SelectItem>
                      <SelectItem value="false">Secondary Factor</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="bobot">Bobot</Label>
                  <Input
                    id="bobot"
                    data-testid="kriteria-bobot-input"
                    type="number"
                    step="0.1"
                    value={formData.bobot}
                    onChange={(e) => setFormData({ ...formData, bobot: e.target.value })}
                    required
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="kriteria-submit-button" className="flex-1">
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

      <Card>
        <CardHeader>
          <CardTitle>Daftar Kriteria</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data kriteria</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Aspek</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Kode</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nama Kriteria</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Tipe</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Bobot</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`kriteria-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4">{getAspekName(item.aspek_id)}</td>
                      <td className="py-3 px-4 font-medium">{item.kode}</td>
                      <td className="py-3 px-4">{item.nama}</td>
                      <td className="py-3 px-4">
                        <span className={`px-2 py-1 text-xs rounded-full ${
                          item.is_core ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700'
                        }`}>
                          {item.is_core ? 'Core' : 'Secondary'}
                        </span>
                      </td>
                      <td className="py-3 px-4 text-gray-600">{item.bobot}</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button size="sm" variant="outline" onClick={() => handleEdit(item)} data-testid={`edit-kriteria-${index}`}>
                              <Edit size={14} />
                            </Button>
                            <Button size="sm" variant="destructive" onClick={() => handleDelete(item.id)} data-testid={`delete-kriteria-${index}`}>
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

export default Kriteria;
