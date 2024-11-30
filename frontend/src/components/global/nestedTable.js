import React, { useState } from 'react';

const NestedTable = ({
  data,
  products = [],
  suppliers = [],
  warehouses = [],
}) => {
  const [expandedRow, setExpandedRow] = useState(null);
  const toggleDetails = (trxPK) => {
    setExpandedRow(expandedRow === trxPK ? null : trxPK);
  };

  return (
    <div className="overflow-x-auto mt-10">
      <table className="w-full text-sm text-left text-gray-500">
        <thead className="text-xs text-gray-700 uppercase bg-gray-300">
          <tr>
            <th className="px-4 py-2 border-b text-left">Transaction No</th>
            <th className="px-4 py-2 border-b text-left">Warehouse ID</th>
            <th className="px-4 py-2 border-b text-left">Transaction Date</th>
            <th className="px-4 py-2 border-b text-left">Supplier ID</th>
            <th className="px-4 py-2 border-b text-left">Notes</th>
            <th className="px-4 py-2 border-b text-left">Details</th>
          </tr>
        </thead>
        <tbody>
          {data.map((trx) => (
            <React.Fragment key={trx.trxOutPK || trx.trxInPK}>
              <tr>
                <td className="px-4 py-2 border-b">
                  {trx.trxOutNo || trx.trxInNo}
                </td>
                <td className="px-4 py-2 border-b">
                  {trx.whsIdf} (
                  {warehouses.find((x) => x.whsPK === trx.whsIdf)?.whsName})
                </td>
                <td className="px-4 py-2 border-b">
                  {new Date(
                    trx.trxOutDate || trx.trxInDate
                  ).toLocaleDateString()}
                </td>
                <td className="px-4 py-2 border-b">
                  {trx.trxOutSuppIdf || trx.trxInSuppIdf} (
                  {
                    suppliers.find(
                      (x) =>
                        String(x.supplierPK) === String(trx.trxOutSuppIdf) ||
                        String(x.supplierPK) === String(trx.trxInSuppIdf)
                    )?.supplierName
                  }
                  )
                </td>
                <td className="px-4 py-2 border-b">
                  {trx.trxOutNotes || trx.trxInNotes}
                </td>
                <td className="px-4 py-2 border-b">
                  <button
                    onClick={() => toggleDetails(trx.trxOutPK || trx.trxInPK)}
                    className="text-indigo-500 hover:underline"
                  >
                    {expandedRow === (trx.trxOutPK || trx.trxInPK)
                      ? 'Hide Details'
                      : 'Show Details'}
                  </button>
                </td>
              </tr>
              {expandedRow === (trx.trxOutPK || trx.trxInPK) && (
                <tr>
                  <td colSpan="6" className="px-4 py-2 border-b bg-gray-100">
                    <table className="min-w-full mt-2 table-auto border-collapse">
                      <thead>
                        <tr>
                          <th className="px-4 py-2 border-b text-left">
                            Product ID
                          </th>
                          <th className="px-4 py-2 border-b text-left">
                            Qty Dus
                          </th>
                          <th className="px-4 py-2 border-b text-left">
                            Qty Pcs
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        {(trx.details || []).map((detail) => (
                          <tr key={detail.trxOutDPK || detail.trxInDPK}>
                            <td className="px-4 py-2 border-b">
                              {detail.trxOutDProductIdf ||
                                detail.trxInDProductIdf}{' '}
                              (
                              {
                                products.find(
                                  (x) =>
                                    String(x.productPK) ===
                                      String(detail.trxOutDProductIdf) ||
                                    String(x.productPK) ===
                                      String(detail.trxInDProductIdf)
                                )?.productName
                              }
                              )
                            </td>
                            <td className="px-4 py-2 border-b">
                              {detail.trxOutDQtyDus || detail.trxInDQtyDus}
                            </td>
                            <td className="px-4 py-2 border-b">
                              {detail.trxOutDQtyPcs || detail.trxInDQtyPcs}
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </td>
                </tr>
              )}
            </React.Fragment>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default NestedTable;
