import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { complianceService, KYCApplication } from '../../services/complianceService';
import { CheckCircleIcon, XCircleIcon, EyeIcon } from '@heroicons/react/24/outline';

export function CompliancePage() {
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [statusFilter, setStatusFilter] = useState('PENDING');
  const [selectedKYC, setSelectedKYC] = useState<KYCApplication | null>(null);
  const queryClient = useQueryClient();

  const { data: kycData, isLoading } = useQuery({
    queryKey: ['kyc', page, limit, statusFilter],
    queryFn: () => complianceService.getKYCList(page, limit, statusFilter),
  });

  const approveMutation = useMutation({
    mutationFn: (kycId: string) => complianceService.approveKYC(kycId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['kyc'] });
    },
  });

  const rejectMutation = useMutation({
    mutationFn: ({ id, reason }: { id: string; reason: string }) => 
      complianceService.rejectKYC(id, reason),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['kyc'] });
    },
  });

  const handleApprove = (id: string) => {
    if (confirm('Are you sure you want to approve this KYC application?')) {
      approveMutation.mutate(id);
    }
  };

  const handleReject = (id: string) => {
    const reason = prompt('Please provide a reason for rejection:');
    if (reason) {
      rejectMutation.mutate({ id, reason });
    }
  };

  const kycs = kycData?.data?.applications || [];
  const total = kycData?.data?.total || 0;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">KYC Compliance</h1>
        <div className="flex space-x-2">
          <select
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="border border-gray-300 rounded-lg px-3 py-2"
          >
            <option value="PENDING">Pending</option>
            <option value="APPROVED">Approved</option>
            <option value="REJECTED">Rejected</option>
            <option value="UNDER_REVIEW">Under Review</option>
          </select>
        </div>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Pending Review</div>
          <div className="text-2xl font-bold text-yellow-600">
            {kycData?.data?.stats?.pending || 0}
          </div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Approved Today</div>
          <div className="text-2xl font-bold text-green-600">
            {kycData?.data?.stats?.approvedToday || 0}
          </div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Rejected Today</div>
          <div className="text-2xl font-bold text-red-600">
            {kycData?.data?.stats?.rejectedToday || 0}
          </div>
        </div>
        <div className="bg-white p-4 rounded-lg shadow">
          <div className="text-sm text-gray-600">Avg Processing Time</div>
          <div className="text-2xl font-bold text-blue-600">
            {kycData?.data?.stats?.avgProcessingTime || 0}h
          </div>
        </div>
      </div>

      {/* KYC List */}
      <div className="bg-white shadow rounded-lg overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Player</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">KYC Level</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Documents</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Submitted</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td colSpan={6} className="px-6 py-4 text-center text-gray-500">Loading...</td>
              </tr>
            ) : kycs.length === 0 ? (
              <tr>
                <td colSpan={6} className="px-6 py-4 text-center text-gray-500">No KYC applications found</td>
              </tr>
            ) : (
              kycs.map((kyc: KYCApplication) => (
                <tr key={kyc.id} className="hover:bg-gray-50">
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="text-sm font-medium text-gray-900">{kyc.playerName}</div>
                    <div className="text-sm text-gray-500">{kyc.playerEmail}</div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {kyc.kycLevel}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {kyc.documents?.length || 0} docs
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {new Date(kyc.submittedAt).toLocaleDateString()}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      kyc.status === 'APPROVED' ? 'bg-green-100 text-green-800' :
                      kyc.status === 'REJECTED' ? 'bg-red-100 text-red-800' :
                      kyc.status === 'UNDER_REVIEW' ? 'bg-blue-100 text-blue-800' :
                      'bg-yellow-100 text-yellow-800'
                    }`}>
                      {kyc.status}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 space-x-2">
                    <button
                      onClick={() => setSelectedKYC(kyc)}
                      className="text-blue-600 hover:text-blue-900"
                    >
                      View
                    </button>
                    {kyc.status === 'PENDING' && (
                      <>
                        <button
                          onClick={() => handleApprove(kyc.id)}
                          className="text-green-600 hover:text-green-900"
                        >
                          Approve
                        </button>
                        <button
                          onClick={() => handleReject(kyc.id)}
                          className="text-red-600 hover:text-red-900"
                        >
                          Reject
                        </button>
                      </>
                    )}
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-between items-center">
        <button
          onClick={() => setPage(p => Math.max(1, p - 1))}
          disabled={page === 1}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Previous
        </button>
        <span className="text-sm text-gray-600">Page {page} - Total: {total}</span>
        <button
          onClick={() => setPage(p => p + 1)}
          disabled={kycs.length < limit}
          className="px-4 py-2 border rounded-lg disabled:opacity-50"
        >
          Next
        </button>
      </div>

      {/* KYC Detail Modal */}
      {selectedKYC && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-2xl w-full max-h-[80vh] overflow-y-auto">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">KYC Application Details</h2>
              <button onClick={() => setSelectedKYC(null)} className="text-gray-500 hover:text-gray-700">
                <XCircleIcon className="h-6 w-6" />
              </button>
            </div>
            <div className="space-y-4">
              <div>
                <h3 className="font-semibold text-gray-700">Player Information</h3>
                <p className="text-sm text-gray-600">Name: {selectedKYC.playerName}</p>
                <p className="text-sm text-gray-600">Email: {selectedKYC.playerEmail}</p>
                <p className="text-sm text-gray-600">Player ID: {selectedKYC.playerId}</p>
              </div>
              <div>
                <h3 className="font-semibold text-gray-700">Documents</h3>
                <ul className="text-sm text-gray-600">
                  {selectedKYC.documents?.map((doc, idx) => (
                    <li key={idx} className="py-1">• {doc.type}: {doc.status}</li>
                  ))}
                </ul>
              </div>
              <div>
                <h3 className="font-semibold text-gray-700">Verification Notes</h3>
                <p className="text-sm text-gray-600">{selectedKYC.notes || 'No notes'}</p>
              </div>
              <div className="flex justify-end space-x-2 pt-4">
                {selectedKYC.status === 'PENDING' && (
                  <>
                    <button
                      onClick={() => handleReject(selectedKYC.id)}
                      className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
                    >
                      Reject
                    </button>
                    <button
                      onClick={() => handleApprove(selectedKYC.id)}
                      className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
                    >
                      Approve
                    </button>
                  </>
                )}
                <button
                  onClick={() => setSelectedKYC(null)}
                  className="px-4 py-2 border rounded-lg"
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
