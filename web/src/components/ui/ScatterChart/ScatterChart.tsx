import {
  ScatterChart,
  CartesianGrid,
  XAxis,
  YAxis,
  ZAxis,
  Tooltip,
  Legend,
  Scatter,
} from "recharts";

export default function MyScatter({ data01 }: { data01: any[] }) {
  return (
    <ScatterChart
      width={730}
      height={250}
      margin={{
        top: 20,
        right: 20,
        bottom: 10,
        left: 10,
      }}
    >
      <CartesianGrid strokeDasharray="3 3" />
      <XAxis dataKey="x" type="number" name="stature" unit="cm" />
      <YAxis dataKey="y" type="number" name="weight" unit="kg" />
      <ZAxis
        dataKey="z"
        type="number"
        range={[64, 144]}
        name="score"
        unit="km"
      />
      <Tooltip cursor={{ strokeDasharray: "3 3" }} />
      <Legend />
      <Scatter name="A school" data={data01} fill="#8884d8" />
    </ScatterChart>
  );
}
