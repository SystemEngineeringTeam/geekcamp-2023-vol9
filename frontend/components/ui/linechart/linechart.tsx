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
import { labels } from "@/const";
import { useHistories } from "@/hooks/useHistories";
import { useDidUpdate } from "@mantine/hooks";
import _ from "lodash";

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip);

export default function Linechart({ roomId }: { roomId: string }) {
  const [histories, { getHistory }] = useHistories();
  const [date, setDate] = useState(getDates(new Date()));
  const [data, setData] = useState({
    labels,
    datasets: [
      {
        data: histories[roomId],
        backgroundColor: "rgb(252, 109, 105)",
      },
    ],
  });

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

  function nextDate() {
    if (date.day === new Date(date.year, date.month, 0).getDate()) {
      const nextDate = new Date(`${date.year}/${date.month + 1}/1`);
      setDate(getDates(nextDate));
    } else {
      const nextDate = new Date(`${date.year}/${date.month}/${date.day + 1}`);
      setDate(getDates(nextDate));
    }
  }

  function prevDate() {
    if (date.day === 1) {
      const prevDate = new Date(
        `${date.year}/${date.month - 1}/${new Date(
          date.year,
          date.month - 1,
          0
        ).getDate()}`
      );
      setDate(getDates(prevDate));
    } else {
      const prevDate = new Date(`${date.year}/${date.month}/${date.day - 1}`);
      setDate(getDates(prevDate));
    }
  }

  function getDates(date: Date) {
    const [year, month, day] = date.toLocaleDateString().split("/");
    return {
      year: parseInt(year),
      month: parseInt(month),
      day: parseInt(day),
    };
  }

  useDidUpdate(() => {
    const tmpDate = `${date.year}-${date.month}-${date.day}`;
    const tmpData = getHistory(roomId, tmpDate);

    if (tmpData !== null) {
      setData((d) => {
        const cloneData = _.cloneDeep(d);
        cloneData.datasets[0].data = tmpData;
        return cloneData;
      });
    }
  }, [date]);

  useDidUpdate(() => {
    setData({
      labels,
      datasets: [
        {
          data: histories[roomId],
          backgroundColor: "rgb(252, 109, 105)",
        },
      ],
    })
  }, [histories]);

  return (
    <div className={styles.linechart}>
      <h2 className={styles.title}>
        {date.year}年{date.month}月{date.day}日
      </h2>

      <div className={styles.linechart_wrapper}>
        <Bar options={options} data={data} />
      </div>

      {/* <div className={styles.arrow} onClick={prevDate}>
        {"<"}
      </div>
      <div className={styles.arrow} onClick={nextDate}>
        {">"}
      </div> */}
    </div>
  );
}
