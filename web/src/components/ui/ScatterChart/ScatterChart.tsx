import {
  ScatterChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Scatter,
} from "recharts";

export default function MyScatter({
  data,
  xkey,
  ykey,
}: {
  data: any[];
  xkey: string;
  ykey: string;
}) {
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
      <XAxis
        dataKey={xkey}
        label={{ value: xkey, position: "insideBottom", offset: 0 }}
        type="number"
      />
      <YAxis
        dataKey={ykey}
        label={{ value: ykey, angle: -90, position: "insideLeft" }}
        type="number"
      />
      <Tooltip cursor={{ strokeDasharray: "3 3" }} />
      <Scatter name={ykey} data={data} fill="#8884d8" />
    </ScatterChart>
  );
}
