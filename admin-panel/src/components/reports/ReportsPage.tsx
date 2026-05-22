import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { reportService, ReportData } from '../../services/reportService';
import { DocumentReportIcon, ChartBarIcon, CurrencyDollarIcon, UserGroupIcon } from '@heroicons/react/24/outline';

export function ReportsPage() {
  const [reportType, setReportType] = useState<'FINANCIAL' | 'PLAYER' | 'GAME' | 'COMPLIANCE'>('FINANCIAL');
  const [dateRange, setDateRange] = useState({ start: '', end: '' });

  const { data: reportsData, isLoading, refetch } = useQuery({
    queryKey: ['reports', reportType, dateRange],
    queryFn: () => reportService.getReport(reportType, dateRange.start, dateRange.end),
    enabled: false,
  });

  const handleGenerateReport = () => {
    if (dateRange.start && dateRange.end) {
      refetch();
    }
  };

  const handleExport = (format: 'PDF' | 'CSV' | 'XLSX') => {
    reportService.exportReport(reportType, dateRange.start, dateRange.end, format);
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">Reports & Analytics</h1>
      </div>

      {/* Report Type Selection */}
      <div className="bg-white p-6 rounded-lg shadow">
        <h2 className="text-lg font-semibold mb-4">Select Report Type</h2>
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
          <button
            onClick={() => setReportType('FINANCIAL')}
            className={`p-4 rounded-lg border-2 flex flex-col items-center ${
              reportType === 'FINANCIAL' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            <CurrencyDollarIcon className="h-8 w-8 text-gray-600 mb-2" />
            <span className="font-medium">Financial</span>
          </button>
          <button
            onClick={() => setReportType('PLAYER')}
            className={`p-4 rounded-lg border-2 flex flex-col items-center ${
              reportType === 'PLAYER' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            <UserGroupIcon className="h-8 w-8 text-gray-600 mb-2" />
            <span className="font-medium">Player</span>
          </button>
          <button
            onClick={() => setReportType('GAME')}
            className={`p-4 rounded-lg border-2 flex flex-col items-center ${
              reportType === 'GAME' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            <ChartBarIcon className="h-8 w-8 text-gray-600 mb-2" />
            <span className="font-medium">Game Performance</span>
          </button>
          <button
            onClick={() => setReportType('COMPLIANCE')}
            className={`p-4 rounded-lg border-2 flex flex-col items-center ${
              reportType === 'COMPLIANCE' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            <DocumentReportIcon className="h-8 w-8 text-gray-600 mb-2" />
            <span className="font-medium">Compliance</span>
          </button>
        </div>
      </div>

      {/* Date Range Selection */}
      <div className="bg-white p-6 rounded-lg shadow">
        <h2 className="text-lg font-semibold mb-4">Date Range</h2>
        <div className="flex flex-wrap gap-4 items-end">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Start Date</label>
            <input
              type="date"
              value={dateRange.start}
              onChange={(e) => setDateRange({ ...dateRange, start: e.target.value })}
              className="border border-gray-300 rounded-md px-3 py-2"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">End Date</label>
            <input
              type="date"
              value={dateRange.end}
              onChange={(e) => setDateRange({ ...dateRange, end: e.target.value })}
              className="border border-gray-300 rounded-md px-3 py-2"
            />
          </div>
          <button
            onClick={handleGenerateReport}
            disabled={!dateRange.start || !dateRange.end}
            className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            Generate Report
          </button>
          {reportsData && (
            <div className="flex gap-2 ml-auto">
              <button
                onClick={() => handleExport('CSV')}
                className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
              >
                Export CSV
              </button>
              <button
                onClick={() => handleExport('XLSX')}
                className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
              >
                Export Excel
              </button>
              <button
                onClick={() => handleExport('PDF')}
                className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
              >
                Export PDF
              </button>
            </div>
          )}
        </div>
      </div>

      {/* Report Results */}
      {isLoading && (
        <div className="bg-white p-6 rounded-lg shadow text-center">
          <p className="text-gray-500">Generating report...</p>
        </div>
      )}

      {reportsData && (
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-lg font-semibold mb-4">{reportType} Report</h2>
          <div className="space-y-4">
            {reportType === 'FINANCIAL' && (
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Total Revenue</p>
                  <p className="text-2xl font-bold text-green-600">
                    ${(reportsData.data as any)?.revenue?.toLocaleString() || '0'}
                  </p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Total Payouts</p>
                  <p className="text-2xl font-bold text-red-600">
                    ${(reportsData.data as any)?.payouts?.toLocaleString() || '0'}
                  </p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Net GGR</p>
                  <p className="text-2xl font-bold text-blue-600">
                    ${(reportsData.data as any)?.ggr?.toLocaleString() || '0'}
                  </p>
                </div>
              </div>
            )}
            {reportType === 'PLAYER' && (
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">New Players</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.newPlayers || 0}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Active Players</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.activePlayers || 0}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Avg Session Time</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.avgSessionTime || 0}m</p>
                </div>
              </div>
            )}
            {reportType === 'GAME' && (
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Total Rounds</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.totalRounds || 0}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Most Popular Game</p>
                  <p className="text-lg font-semibold">{(reportsData.data as any)?.topGame || 'N/A'}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">Avg RTP</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.avgRtp || 0}%</p>
                </div>
              </div>
            )}
            {reportType === 'COMPLIANCE' && (
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">KYC Pending</p>
                  <p className="text-2xl font-bold text-yellow-600">{(reportsData.data as any)?.kycPending || 0}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">AML Alerts</p>
                  <p className="text-2xl font-bold text-red-600">{(reportsData.data as any)?.amlAlerts || 0}</p>
                </div>
                <div className="p-4 bg-gray-50 rounded-lg">
                  <p className="text-sm text-gray-600">SAR Filings</p>
                  <p className="text-2xl font-bold">{(reportsData.data as any)?.sarFilings || 0}</p>
                </div>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
