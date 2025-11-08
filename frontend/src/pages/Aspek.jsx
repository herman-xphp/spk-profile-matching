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

const Aspek = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({ id: '', nama: '', persentase: '' });

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/aspek`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data aspek');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { nama: formData.nama, persentase: parseFloat(formData.persentase) };
    try {
      if (editMode) {
        await axios.put(`${API}/aspek/${formData.id}`, payload);
        toast.success('Aspek berhasil diperbarui');
      } else {
        await axios.post(`${API}/aspek`, payload);
        toast.success('Aspek berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      fetchData();
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus aspek ini?')) return;
    try {
      await axios.delete(`${API}/aspek/${id}`);
      toast.success('Aspek berhasil dihapus');
      fetchData();
    } catch (error) {
      toast.error('Gagal menghapus aspek');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', nama: '', persentase: '' });
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
  const totalPersentase = data.reduce((sum, item) => sum + item.persentase, 0);

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Data Aspek</h1>
          <p className="text-gray-500 mt-1">Kelola aspek penilaian dengan persentase bobot</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-aspek-button">
                <Plus size={16} className="mr-2" />
                Tambah Aspek
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Aspek' : 'Tambah Aspek'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="nama">Nama Aspek</Label>
                  <Input
                    id="nama"
                    data-testid="aspek-nama-input"
                    value={formData.nama}
                    onChange={(e) => setFormData({ ...formData, nama: e.target.value })}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="persentase">Persentase (%)</Label>
                  <Input
                    id="persentase"
                    data-testid="aspek-persentase-input"
                    type="number"
                    step="0.1"
                    min="0"
                    max="100"
                    value={formData.persentase}
                    onChange={(e) => setFormData({ ...formData, persentase: e.target.value })}
                    required
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="aspek-submit-button" className="flex-1">
                    {editMode ? 'Update' : 'Simpan'}
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => setOpen(false)}
                    className="flex-1"
                  >
                    Batal
                  </Button>
                </div>
              </form>
            </DialogContent>
          </Dialog>
        )}
      </div>

      {totalPersentase !== 100 && data.length > 0 && (
        <div className="mb-4 p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
          <p className="text-yellow-800 text-sm">
            <strong>Peringatan:</strong> Total persentase saat ini adalah {totalPersentase}%. Harus 100%.
          </p>
        </div>
      )}

      <Card>
        <CardHeader>
          <CardTitle>Daftar Aspek</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data aspek</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nama Aspek</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Persentase</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`aspek-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4 font-medium">{item.nama}</td>
                      <td className="py-3 px-4 text-gray-600">{item.persentase}%</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button
                              size="sm"
                              variant="outline"
                              onClick={() => handleEdit(item)}
                              data-testid={`edit-aspek-${index}`}
                            >
                              <Edit size={14} />
                            </Button>
                            <Button
                              size="sm"
                              variant="destructive"
                              onClick={() => handleDelete(item.id)}
                              data-testid={`delete-aspek-${index}`}
                            >
                              <Trash2 size={14} />
                            </Button>
                          </div>
                        </td>
                      )}
                    </tr>
                  ))}
                </tbody>
                <tfoot>
                  <tr className="border-t-2 font-semibold">
                    <td colSpan="2" className="py-3 px-4 text-right">Total:</td>
                    <td className="py-3 px-4">{totalPersentase}%</td>
                    {isAdmin && <td></td>}
                  </tr>
                </tfoot>
              </table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default Aspek;
