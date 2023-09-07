import styles from "./linechart.module.scss";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Tooltip,
} from "chart.js";
import { Bar } from "react-chartjs-2";
import { faker } from "@faker-js/faker";
import { _DeepPartialObject } from "chart.js/dist/types/utils";
import { useState } from "react";

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip);

export default function Linechart() {
  const options = {
    maintainAspectRatio: false,
    animation: {
      duration: 1000,
    },
    scales: {
      y: {
        max: undefined,
      },
      x: {
        ticks: {
          maxRotation: 0, // ラベルを横に表示
          minRotation: 0, // ラベルを横に表示
        },
      },
    },
    plugins: {
      legend: {
        position: "top" as const,
      },
      title: {
        display: false,
      },
    },
  };
  const labels = Array.from({ length: 24 }, (_, index) => `${index}時`);
  const data = {
    labels,
    datasets: [
      {
        data: labels.map(() => faker.number.int({ min: 0, max: 1000 })),
        backgroundColor: "rgb(252, 109, 105)",
      },
    ],
  };

  return (
    <div className={styles.linechart_wrapper}>
      <Bar options={options} data={data} updateMode={"resize"} />
    </div>
  );
}
