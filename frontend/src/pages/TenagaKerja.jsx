import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { API, AuthContext } from '../App';
import { Button } from '../components/ui/button';
import { Input } from '../components/ui/input';
import { Label } from '../components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '../components/ui/dialog';
import { toast } from 'sonner';
import { Plus, Edit, Trash2 } from 'lucide-react';

const TenagaKerja = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({
    id: '', nik: '', nama: '', tgl_lahir: '', alamat: '', telepon: ''
  });

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/tenaga-kerja`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data tenaga kerja');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      nik: formData.nik,
      nama: formData.nama,
      tgl_lahir: formData.tgl_lahir,
      alamat: formData.alamat,
      telepon: formData.telepon
    };
    try {
      if (editMode) {
        await axios.put(`${API}/tenaga-kerja/${formData.id}`, payload);
        toast.success('Data tenaga kerja berhasil diperbarui');
      } else {
        await axios.post(`${API}/tenaga-kerja`, payload);
        toast.success('Data tenaga kerja berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      fetchData();
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus tenaga kerja ini?')) return;
    try {
      await axios.delete(`${API}/tenaga-kerja/${id}`);
      toast.success('Data tenaga kerja berhasil dihapus');
      fetchData();
    } catch (error) {
      toast.error('Gagal menghapus data tenaga kerja');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', nik: '', nama: '', tgl_lahir: '', alamat: '', telepon: '' });
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

  const isAdmin = user?.role === 'admin';

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Data Tenaga Kerja</h1>
          <p className="text-gray-500 mt-1">Kelola data tenaga kerja pabrik</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-tenaga-kerja-button">
                <Plus size={16} className="mr-2" />
                Tambah Tenaga Kerja
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Tenaga Kerja' : 'Tambah Tenaga Kerja'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="nik">NIK</Label>
                  <Input
                    id="nik"
                    data-testid="tk-nik-input"
                    value={formData.nik}
                    onChange={(e) => setFormData({ ...formData, nik: e.target.value })}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="nama">Nama Lengkap</Label>
                  <Input
                    id="nama"
                    data-testid="tk-nama-input"
                    value={formData.nama}
                    onChange={(e) => setFormData({ ...formData, nama: e.target.value })}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="tgl_lahir">Tanggal Lahir</Label>
                  <Input
                    id="tgl_lahir"
                    data-testid="tk-tgl-lahir-input"
                    type="date"
                    value={formData.tgl_lahir || ''}
                    onChange={(e) => setFormData({ ...formData, tgl_lahir: e.target.value })}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="alamat">Alamat</Label>
                  <Input
                    id="alamat"
                    data-testid="tk-alamat-input"
                    value={formData.alamat || ''}
                    onChange={(e) => setFormData({ ...formData, alamat: e.target.value })}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="telepon">Telepon</Label>
                  <Input
                    id="telepon"
                    data-testid="tk-telepon-input"
                    value={formData.telepon || ''}
                    onChange={(e) => setFormData({ ...formData, telepon: e.target.value })}
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="tk-submit-button" className="flex-1">
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
          <CardTitle>Daftar Tenaga Kerja</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data tenaga kerja</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">NIK</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nama</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Tanggal Lahir</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Telepon</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`tk-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4 font-medium">{item.nik}</td>
                      <td className="py-3 px-4">{item.nama}</td>
                      <td className="py-3 px-4 text-gray-600">{item.tgl_lahir || '-'}</td>
                      <td className="py-3 px-4 text-gray-600">{item.telepon || '-'}</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button size="sm" variant="outline" onClick={() => handleEdit(item)} data-testid={`edit-tk-${index}`}>
                              <Edit size={14} />
                            </Button>
                            <Button size="sm" variant="destructive" onClick={() => handleDelete(item.id)} data-testid={`delete-tk-${index}`}>
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

export default TenagaKerja;
